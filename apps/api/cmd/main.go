// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/your-username/dilocash-oss/apps/api/internal/infra/storage/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env")
	}

	// 2. Initialize Database Connection
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	db, err := postgres.NewConnection(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 3. Initialize Repositories
	// We inject the DB connection into our repo implementation
	_ = postgres.NewTransactionRepository(db)
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

	// Enable Reflection for easier debugging (e.g., using Postman or gRPCui)
	reflection.Register(grpcServer)

	// 5. Graceful Shutdown Handling
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
	grpcServer.GracefulStop()
	log.Println("Server stopped.")
}
