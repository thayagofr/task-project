package middleware

import (
	"net/http"
	"strings"
	"task-api/internal/adapters/in/rest/consts"
	"task-api/internal/ports/out"
)

type JWTMiddleware struct {
	jwtManager out.TokenManager
}

func NewJWTMiddleware(jwtManager out.TokenManager) *JWTMiddleware {
	return &JWTMiddleware{jwtManager: jwtManager}
}

func (middleware *JWTMiddleware) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(consts.Authorization)
		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		if err := middleware.jwtManager.Validate(token); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
