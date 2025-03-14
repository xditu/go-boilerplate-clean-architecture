package interface_adapters

import (
	"encoding/json"
	"net/http"
)

type HelloWorldAdapter struct {
}

func (uc *HelloWorldAdapter) HelloWorldAdapter(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{Message: "Hello, World!"})
	if err != nil {
		return
	}
}

func NewHelloWorldAdapter() *HelloWorldAdapter {
	return &HelloWorldAdapter{}
}
