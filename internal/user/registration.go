package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"board/pkg/utils"
)

type Registration struct {
	rep *Repository
}

func NewRegistration(rep *Repository) *Registration {
	if rep == nil {
		panic("nil repository")
	}

	return &Registration{rep: rep}
}

func (r *Registration) Register(auth Auth) error {
	if auth.Password != auth.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	err := r.rep.CreateUser(User{
		ID:       uuid.New(),
		Nickname: auth.Nickname,
		Avatar:   auth.Avatar,
		Hash:     utils.Hash(auth.Password),
	})
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
