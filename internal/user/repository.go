package user

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(user User) error
	SearchUserByNickName(nickName string) (*User, error)
	FindUser(id uuid.UUID) (*User, error)
}

type memoryRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*User
}

func NewRepository() UserRepository {
	return &memoryRepository{
		mu:    sync.RWMutex{},
		users: make(map[uuid.UUID]*User),
	}
}

func (r *memoryRepository) CreateUser(u User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[u.ID] = &u
	return nil
}

func (r *memoryRepository) SearchUserByNickName(nickName string) (*User, error) {
	for _, u := range r.users {
		if u.Nickname == nickName {
			return u, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (r *memoryRepository) FindUser(id uuid.UUID) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

// postgresRepository — реализация UserRepository для PostgreSQL.
type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) UserRepository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) CreateUser(u User) error {
	_, err := r.pool.Exec(context.Background(),
		`INSERT INTO users (id, nickname, avatar, hash) VALUES ($1, $2, $3, $4)`,
		u.ID, u.Nickname, u.Avatar, u.Hash)
	return err
}

func (r *postgresRepository) SearchUserByNickName(nickName string) (*User, error) {
	var u User
	err := r.pool.QueryRow(context.Background(),
		`SELECT id, nickname, avatar, hash FROM users WHERE nickname = $1`, nickName).
		Scan(&u.ID, &u.Nickname, &u.Avatar, &u.Hash)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &u, nil
}

func (r *postgresRepository) FindUser(id uuid.UUID) (*User, error) {
	var u User
	err := r.pool.QueryRow(context.Background(),
		`SELECT id, nickname, avatar, hash FROM users WHERE id = $1`, id).
		Scan(&u.ID, &u.Nickname, &u.Avatar, &u.Hash)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &u, nil
}
