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

	db "github.com/dilocash/dilocash-oss/internal/generated/db/postgres"
	"github.com/dilocash/dilocash-oss/internal/infra/health"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}

	// 2. Initialize Database Connection (pgxpool for sqlc compatibility)
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	// 3. Initialize SQLC Querier
	_ = db.New(pool)
	// (Note: In future steps, we'll inject this into our Service/Use-Case layer)

	// 4. Setup gRPC Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	// Enable Reflection for easier debugging
	reflection.Register(grpcServer)

	// Register Health Service
	healthManager := health.NewManager(pool)
	healthManager.Register(grpcServer)

	// 5. Database Health Monitor
	healthIntervalStr := os.Getenv("DB_HEALTH_CHECK_INTERVAL")
	healthInterval := 60 * time.Second
	if healthIntervalStr != "" {
		if val, err := strconv.Atoi(healthIntervalStr); err == nil {
			healthInterval = time.Duration(val) * time.Second
		}
	}

	healthCtx, cancelHealth := context.WithCancel(context.Background())
	defer cancelHealth()

	go healthManager.Monitor(healthCtx, healthInterval)

	// 6. Graceful Shutdown Handling
	go func() {
		log.Printf("🚀 Dilocash-OSS API starting on port %s", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	cancelHealth() // Stop the health monitor

	// Create a context that will timeout after 30 seconds
	// to ensure the process eventually exits if draining hangs.
	stopped := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		log.Println("Drained all connections successfully.")
	case <-time.After(30 * time.Second):
		log.Println("Shutdown timed out, forcing stop...")
		grpcServer.Stop()
	}

	log.Println("Server stopped.")
}
