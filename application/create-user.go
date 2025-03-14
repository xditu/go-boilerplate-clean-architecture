package application

import (
	"go-boilerplate-clean-architecture/enterprise"
)

type CreateUserUseCase struct {
	// UserRepo is an interface that defines the methods required for user repository operations.
	// It is used to interact with the underlying data storage for user-related actions.
	UserRepo enterprise.UserRepositories
}

func (uc *CreateUserUseCase) Execute(name, lastname, username string) (*enterprise.User, error) {
	user := enterprise.User{
		Name:     name,
		Lastname: lastname,
		Username: username,
	}
	if err := user.ValidateName(); err != nil {
		return nil, err
	}
	err := uc.UserRepo.Add(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewCreateUser(repo enterprise.UserRepositories) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepo: repo}
}
