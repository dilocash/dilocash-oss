// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"strconv"

	"net/http"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/validate"
	db "github.com/dilocash/dilocash-oss/apps/api/internal/generated/db/postgres"
	"github.com/dilocash/dilocash-oss/apps/api/internal/generated/transport/dilocash/v1/v1connect"
	"github.com/dilocash/dilocash-oss/apps/api/internal/infra/health"
	"github.com/dilocash/dilocash-oss/apps/api/internal/middleware"
	"github.com/dilocash/dilocash-oss/apps/api/internal/services/sync"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://localhost:8080"},
		AllowedMethods:   connectcors.AllowedMethods(),
		AllowedHeaders:   append(connectcors.AllowedHeaders(), "Authorization"),
		ExposedHeaders:   connectcors.ExposedHeaders(),
		AllowCredentials: true,
	})
	return middleware.Handler(h)
}

// registerAllServices registers all gRPC services from v1
func registerAllServices(ctx context.Context, mux *http.ServeMux, grpcServer *grpc.Server, pool *pgxpool.Pool) {

	log.Println("Configure auth server")
	supabaseAuth := configureAuthServer(ctx)

	log.Println("Registering gRPC services...")
	// sync server
	syncServer := sync.NewSyncServer(pool)
	path, handler := v1connect.NewSyncServiceHandler(
		syncServer,
		connect.WithInterceptors(
			middleware.NewAuthInterceptor(&supabaseAuth),
			validate.NewInterceptor()),
	)
	// Apply the middleware to protected routes
	mux.Handle(path, withCORS(handler))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 1. Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}

	// 2. Initialize Database Connection (pgxpool for sqlc compatibility)
	pool := initDB()
	defer pool.Close()

	// 4. Setup gRPC Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create gRPC server with keepalive settings
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
			MaxConnectionAge:  10 * time.Minute,
		}),
	)

	// Enable Reflection for easier debugging
	reflection.Register(grpcServer)
	cancelHealth := configureHealthCheck(pool, grpcServer)
	defer cancelHealth()
	// Setup HTTP/2 Clear Text (h2c) server to handle both gRPC and HTTP calls
	go func() {
		log.Printf("🚀 Dilocash-OSS API starting on port %s", port)

		// Create a custom HTTP server that supports both HTTP and gRPC over the same port
		mux := http.NewServeMux()

		// Register all services
		registerAllServices(ctx, mux, grpcServer, pool)

		p := new(http.Protocols)
		p.SetHTTP1(true)
		p.SetUnencryptedHTTP2(true)

		// Use h2c to allow both HTTP and gRPC on the same port
		h2cServer := http.Server{
			Addr:      ":" + port,
			Handler:   withCORS(mux),
			Protocols: p,
		}

		// Set up the main HTTP handler to route gRPC calls to the gRPC server
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Check if this is a gRPC request
			if r.ProtoMajor == 2 && r.Header.Get("content-type") == "application/grpc" {
				grpcServer.ServeHTTP(w, r)
			} else {
				// Handle regular HTTP requests
				http.DefaultServeMux.ServeHTTP(w, r)
			}
		})

		// Start the server
		if err := h2cServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	cancelHealth() // Stop the health monitor

	// Create a context that will timeout after 30 seconds
	// to ensure the process eventually exits if draining hangs.
	_, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Stop the HTTP server gracefully
	// Note: We need to access h2cServer from the goroutine scope
	// For now, we'll just gracefully stop the gRPC server
	grpcServer.GracefulStop()
	log.Println("Server stopped.")
}

func initDB() *pgxpool.Pool {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	//
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize SQLC Querier
	_ = db.New(pool)

	return pool
}

func configureHealthCheck(pool *pgxpool.Pool, grpcServer *grpc.Server) context.CancelFunc {
	// Set up Database Health Monitor
	healthIntervalStr := os.Getenv("DB_HEALTH_CHECK_INTERVAL")
	healthInterval := 60 * time.Second
	if healthIntervalStr != "" {
		if val, err := strconv.Atoi(healthIntervalStr); err == nil {
			healthInterval = time.Duration(val) * time.Second
		}
	}

	healthCtx, cancelHealth := context.WithCancel(context.Background())

	// Create health manager instance for monitoring
	healthManager := health.NewManager(pool)
	go healthManager.Monitor(healthCtx, healthInterval) // Register health service
	healthManager.Register(grpcServer)

	return cancelHealth
}

func configureAuthServer(ctx context.Context) middleware.SupabaseAuth {
	supabaseServer := os.Getenv("SUPABASE_SERVER")
	if supabaseServer == "" {
		log.Fatalf("SUPABASE_SERVER environment variable not set")
	}

	supabaseAuth, err := middleware.NewSupabaseAuth(ctx, supabaseServer)
	if err != nil {
		log.Fatalf("NewSupabaseAuth init failed: %v", err)
	}
	return *supabaseAuth
}
