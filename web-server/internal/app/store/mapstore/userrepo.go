package mapstore

import (
	"errors"

	"github.com/artemxgod/11-go-projects/web-server/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepository) Find(ID int) (*model.User, error) {
	if _, ok := r.users[ID]; ok {
		return r.users[ID], nil
	}
	return &model.User{}, errors.New("record not found")
}