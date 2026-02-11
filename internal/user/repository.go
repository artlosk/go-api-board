package user

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Repository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*User
}

func NewRepository() *Repository {
	return &Repository{
		mu:    sync.RWMutex{},
		users: make(map[uuid.UUID]*User),
	}
}

func (r *Repository) CreateUser(user User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = &user

	return nil
}

func (r *Repository) SearchUserByNickName(nickName string) (*User, error) {
	for _, user := range r.users {
		if user.Nickname == nickName {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (r *Repository) FindUser(id uuid.UUID) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}
