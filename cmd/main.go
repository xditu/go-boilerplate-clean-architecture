package main

import (
	"go-boilerplate-clean-architecture/frameworks-drivers"
	"go-boilerplate-clean-architecture/frameworks-drivers/controllers"
)

func main() {
	userRepository := frameworks_drivers.NewMemUserRepository()
	uController := controllers.NewUserController(userRepository)
	hwController := controllers.NewHelloWorldController()

	s := frameworks_drivers.NewServer(8080, hwController, uController)
	s.Serve()
}
