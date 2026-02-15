package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"board/internal/board"
	"board/internal/database"
	"board/internal/server"
	"board/internal/user"
	"board/pkg/jwt"
)

func main() {
	_ = godotenv.Load() // игнорируем ошибку — .env может отсутствовать

	srv := server.NewServer()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "qwerty"
	}

	var userRep user.UserRepository
	var boardRep board.BoardRepository

	dsn := os.Getenv("DATABASE_URL")
	if dsn != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		pool, err := database.OpenPool(ctx, dsn)
		if err != nil {
			panic("postgres: " + err.Error())
		}
		defer pool.Close()
		userRep = user.NewPostgresRepository(pool)
		boardRep = board.NewPostgresRepository(pool)
	} else {
		userRep = user.NewRepository()
		boardRep = board.NewRepository()
	}

	regService := user.NewRegistration(userRep)
	j := jwt.NewJwt(secret)
	loginService := user.NewLoginService(userRep, j)
	b := board.NewBoard(boardRep)

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
