// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package health

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Manager handles the gRPC health service and background monitoring.
type Manager struct {
	server *health.Server
	pool   *pgxpool.Pool
}

// NewManager creates a new health manager.
func NewManager(pool *pgxpool.Pool) *Manager {
	return &Manager{
		server: health.NewServer(),
		pool:   pool,
	}
}

// Register attaches the health service to the provided gRPC server.
func (m *Manager) Register(s *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(s, m.server)
	// Set initial status
	m.server.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
}

// Monitor starts a background goroutine to check database connectivity.
func (m *Manager) Monitor(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := m.pool.Ping(ctx); err != nil {
				log.Printf("⚠️  Database health check failed: %v", err)
				m.server.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
			} else {
				m.server.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
			}
		}
	}
}
