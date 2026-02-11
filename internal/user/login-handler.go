package user

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"board/pkg/jwt"
)

func MakeLoginHandler(jwt *jwt.Jwt, loginService *LoginService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		var l Login

		if err := json.Unmarshal(body, &l); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		s, err := loginService.Login(l.Login, l.Password)
		if err != nil {
			slog.Error("error login", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write([]byte(s))
		if err != nil {
			return
		}
	}
}
