package main

import (
	"net/http"
	"os"

	"board/internal/board"
	"board/internal/server"
	"board/internal/user"
	"board/pkg/jwt"
)

func main() {
	srv := server.NewServer()

	secret := "qwerty"
	if len(os.Args) > 1 {
		secret = os.Args[1]
	}

	userRep := user.NewRepository()

	regService := user.NewRegistration(userRep)

	j := jwt.NewJwt(secret)

	loginService := user.NewLoginService(userRep, j)

	b := board.NewBoard(board.NewRepository())

	srv.AddHandler(http.MethodPost, "/", user.MakeMiddlewareAuth(j, loginService,
		board.CreateAnnouncementHandler(b),
	))

	srv.AddHandler(http.MethodPut, "/", user.MakeMiddlewareAuth(j, loginService,
		board.UpdateAnnouncementHandler(b),
	))

	srv.AddHandler(http.MethodDelete, "/", user.MakeMiddlewareAuth(j, loginService,
		board.DeleteAnnouncementHandler(b),
	))

	srv.AddHandler(http.MethodGet, "/", board.ListHandler(b))

	srv.AddHandler(http.MethodPost, "/reg", user.MakeRegHandler(regService))

	srv.AddHandler(http.MethodPost, "/login", user.MakeLoginHandler(j, loginService))

	err := srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
