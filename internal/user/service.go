package user

import (
	"fmt"
	"log/slog"

	"board/pkg/jwt"
	"board/pkg/utils"
)

type LoginService struct {
	rep *Repository
	jwt *jwt.Jwt
}

func NewLoginService(rep *Repository, jwt *jwt.Jwt) *LoginService {
	return &LoginService{rep: rep, jwt: jwt}
}

func (l *LoginService) Login(nickname, password string) (string, error) {
	user, err := l.rep.SearchUserByNickName(nickname)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	if user.Hash != utils.Hash(password) {
		return "", fmt.Errorf("invalid password")
	}

	signingString, err := l.jwt.SigningString(user.ID)
	if err != nil {
		slog.Error("error ctrl", slog.String("error", err.Error()))

		return "", fmt.Errorf("jwt signing error")
	}

	return signingString, nil
}
