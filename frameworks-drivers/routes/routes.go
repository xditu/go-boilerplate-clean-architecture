package routes

import (
	"go-boilerplate-clean-architecture/frameworks-drivers/controllers"
	"net/http"
)

func SetupUserRoutes(userController *controllers.UserController) {
	http.HandleFunc("/users", userController.CreateUserController)
}

func SetupHelloWorldRoutes(helloWorldController *controllers.HelloWorldController) {
	http.HandleFunc("/hello-world", helloWorldController.HelloWorldController)
}
