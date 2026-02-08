package middleware

import (
	"context"
	"fmt"
	"log"

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

	return &SupabaseAuth{
		keyfunc: k.Keyfunc,
		issuer:  supabaseServer + "/auth/v1",
	}, nil
}

func (s *SupabaseAuth) Validate(tokenString string) (*jwt.Token, error) {
	// 1. Define your Supabase JWKS URL

	// 3. Parse and Validate
	token, err := jwt.Parse(tokenString, s.keyfunc,
		jwt.WithIssuer(s.issuer),
		jwt.WithAudience("authenticated"), // Default for Supabase user tokens
	)

	if err != nil {
		log.Fatalf("Token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Authenticated user ID: %v\n", claims["sub"])
	}
	return token, nil
}
