package middleware

import (
	"context"
	"go_api_cimol/helpers"
	"net/http"

	"golang.org/x/time/rate"
)

type KeyUserInfo string

var limiter = rate.NewLimiter(5, 1)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accesToken := r.Header.Get("Authorization")

		if accesToken == "" {
			helpers.Response(w, 401, "unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(accesToken)
		if err != nil {
			helpers.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUserInfo("userinfo"), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
