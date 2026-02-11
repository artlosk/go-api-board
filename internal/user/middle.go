package user

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"board/pkg/jwt"
)

func MakeMiddlewareAuth(jwt *jwt.Jwt, loginService *LoginService, next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		part := strings.Split(authorization, " ")
		if len(part) != 2 || part[0] != "Bearer" {
			return
		}

		id, err := jwt.GetID(part[1])
		if err != nil {
			slog.Error("jwt get id err", slog.String("err", err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)

		r, err = http.NewRequestWithContext(ctx, r.Method, r.RequestURI, r.Body)
		if err != nil {
			return
		}

		next(w, r)
	}
}
