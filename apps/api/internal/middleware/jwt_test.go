// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	v1 "github.com/dilocash/dilocash-oss/apps/api/internal/generated/transport/dilocash/v1"
	"github.com/dilocash/dilocash-oss/apps/api/internal/generated/transport/dilocash/v1/v1connect"
	"github.com/dilocash/dilocash-oss/apps/api/internal/services/sync"
)

func TestAuthInterceptor(t *testing.T) {
	ctx := context.Background()

	supabaseAuth, err := NewSupabaseAuth(ctx, "")
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	syncServer := sync.NewSyncServer(nil, nil)
	path, handler := v1connect.NewSyncServiceHandler(
		syncServer,
		connect.WithInterceptors(
			NewAuthInterceptor(supabaseAuth),
			validate.NewInterceptor()),
	)
	// Apply the middleware to protected routes
	mux.Handle(path, handler)

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client := v1connect.NewSyncServiceClient(
		http.DefaultClient,
		srv.URL,
	)

	t.Run("Must fail if no token provided", func(t *testing.T) {
		_, err := client.PullChanges(context.Background(), &v1.PullChangesRequest{})
		if connect.CodeOf(err) != connect.CodeUnauthenticated {
			t.Errorf("expected Unauthenticated, got %v", err)
		}
	})
}
