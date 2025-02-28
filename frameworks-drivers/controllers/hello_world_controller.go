package controllers

import (
	"encoding/json"
	"net/http"
)

type HelloWorldController struct {
}

func (uc *HelloWorldController) HelloWorldController(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{Message: "Hello, World!"})
	if err != nil {
		return
	}
}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{}
}
