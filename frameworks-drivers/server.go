package frameworks_drivers

import (
	"fmt"
	"go-boilerplate-clean-architecture/frameworks-drivers/controllers"
	"go-boilerplate-clean-architecture/frameworks-drivers/routes"
	"net/http"
)

type Server struct {
	Port int `json:"port"`
}

func (o *Server) Serve() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Listening on port :8080")
	}
}

func NewServer(port int,
	hwController *controllers.HelloWorldController,
	uController *controllers.UserController) *Server {
	r := &Server{Port: port}
	routes.SetupHelloWorldRoutes(hwController)
	routes.SetupUserRoutes(uController)
	return r
}
