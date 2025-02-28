package database

import (
	"errors"
	"go-boilerplate-clean-architecture/enterprise"
)

type MemUserRepository struct {
	m map[string]enterprise.User
}

func (o *MemUserRepository) Add(u *enterprise.User) error {
	o.m[u.Username] = *u
	return nil
}

func (o *MemUserRepository) Remove(u enterprise.User) error {
	delete(o.m, u.Username)
	return nil
}

func (o *MemUserRepository) Update(u enterprise.User) error {
	return nil
}

func (o *MemUserRepository) Get(username string) (enterprise.User, error) {
	if r, exists := o.m[username]; exists {
		return r, nil
	}
	return enterprise.User{}, errors.New("user not found")
}

func NewMemUserRepository() *MemUserRepository {
	return &MemUserRepository{m: map[string]enterprise.User{}}
}
