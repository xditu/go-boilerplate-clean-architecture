package database

import (
	"errors"
	"go-boilerplate-clean-architecture/enterprise"
)

type MySQLUserRepository struct {
	//host, port, username, password, etc
}

func (o *MySQLUserRepository) Add(u *enterprise.User) error {
	return errors.New("not implemented")
}

func (o *MySQLUserRepository) Remove(u enterprise.User) error {
	return errors.New("not implemented")
}

func (o *MySQLUserRepository) Update(u enterprise.User) error {
	return errors.New("not implemented")
}

func (o *MySQLUserRepository) Get(username string) (enterprise.User, error) {
	return enterprise.User{}, errors.New("not implemented")
}

func NewMySQLUserRepository() *MySQLUserRepository {
	return &MySQLUserRepository{}
}
