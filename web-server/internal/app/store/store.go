package store

import "github.com/artemxgod/11-go-projects/web-server/internal/app/model"

type Store interface {
	User() UserRepository
}

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}