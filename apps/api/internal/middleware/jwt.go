// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package middleware

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type SupabaseAuth struct {
	keyfunc jwt.Keyfunc
	issuer  string
}

func NewSupabaseAuth(ctx context.Context, supabaseServer string) (*SupabaseAuth, error) {
	jwksURL := supabaseServer + "/auth/v1/.well-known/jwks.json"

	// Initialize with background refresh.
	// Default options include automatic periodic refreshes.
	k, err := keyfunc.NewDefaultCtx(ctx, []string{jwksURL})
	if err != nil {
		return nil, err
	}
	log.Printf("Auth Issuer %s", supabaseServer+"/auth/v1")
	return &SupabaseAuth{
		keyfunc: k.Keyfunc,
		issuer:  supabaseServer + "/auth/v1",
	}, nil
}

func NewAuthInterceptor(auth *SupabaseAuth) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// 1. Extract Bearer token
			authHeader := req.Header().Get("Authorization")
			if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing token"))
			}

			// 2. Validate using cached JWKS
			token, err := auth.Validate(authHeader[7:])
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			// 3. Add user ID to context for downstream use
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				ctx = context.WithValue(ctx, "user_id", claims["sub"])
			}

			return next(ctx, req)
		}
	}
}

func (s *SupabaseAuth) Validate(tokenString string) (*jwt.Token, error) {
	// 1. Define your Supabase JWKS URL

	// 3. Parse and Validate
	token, err := jwt.Parse(tokenString, s.keyfunc,
		jwt.WithIssuer(s.issuer),
		jwt.WithAudience("authenticated"), // Default for Supabase user tokens
	)

	if err != nil {
		log.Printf("Token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("Authenticated user ID: %v\n", claims["sub"])
	}
	return token, nil
}
