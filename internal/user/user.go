package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Nickname string    `json:"nickname"`
	Avatar   string    `json:"avatar"`
	Hash     string    `json:"hash"`
}
