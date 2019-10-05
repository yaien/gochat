package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stretchr/objx"
)

// Auth middleware injects 'user' in request context
func Auth(next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		data := objx.MustFromBase64(cookie.Value)
		user := &User{
			Name:     data.Get("name").Str(""),
			Email:    data.Get("email").Str(""),
			Username: data.Get("username").Str(""),
		}
		fmt.Println(user)
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(middleware)
}
