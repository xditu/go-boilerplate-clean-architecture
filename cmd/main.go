package main

import (
	"go-boilerplate-clean-architecture/frameworks-drivers"
	"go-boilerplate-clean-architecture/frameworks-drivers/controllers"
	"go-boilerplate-clean-architecture/frameworks-drivers/database"
)

func main() {
	userRepository := database.NewMemUserRepository()
	uController := controllers.NewUserController(userRepository)
	hwController := controllers.NewHelloWorldController()

	s := frameworks_drivers.NewServer(8081, hwController, uController)
	s.Serve()
}
