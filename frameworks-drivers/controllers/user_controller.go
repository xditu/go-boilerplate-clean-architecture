package controllers

import (
	"encoding/json"
	"go-boilerplate-clean-architecture/application"
	"go-boilerplate-clean-architecture/enterprise"
	"net/http"
)

type UserController struct {
	UserCase *application.CreateUserUseCase
}

func (uc *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{Message: "Hello, World!"})
	if err != nil {
		return
	}
}

func NewUserController(repository enterprise.UserRepository) *UserController {
	return &UserController{UserCase: application.NewCreateUser(repository)}
}
