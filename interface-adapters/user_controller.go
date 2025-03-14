package interface_adapters

import (
	"encoding/json"
	"go-boilerplate-clean-architecture/application"
	"go-boilerplate-clean-architecture/enterprise"
	"net/http"
)

type UserAdapter struct {
	UserCase *application.CreateUserUseCase
}

func (uc *UserAdapter) CreateUserAdapter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{Message: "Hello, World!"})
	if err != nil {
		return
	}
}

func NewUserAdapter(repository enterprise.UserRepositories) *UserAdapter {
	return &UserAdapter{UserCase: application.NewCreateUser(repository)}
}
