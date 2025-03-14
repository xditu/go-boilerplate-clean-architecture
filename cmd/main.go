package main

import (
	"fmt"
	"go-boilerplate-clean-architecture/frameworks-drivers/controllers"
	"go-boilerplate-clean-architecture/frameworks-drivers/database"
	interface_adapters "go-boilerplate-clean-architecture/interface-adapters"
	"sync"
)

func main() {
	userRepository := database.NewMemUserRepository()
	uController := interface_adapters.NewUserAdapter(userRepository)
	hwController := interface_adapters.NewHelloWorldAdapter()
	s := controllers.NewServerHTTP(8081, hwController, uController)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		s.ServeHTTP()
	}()

	go func() {
		defer wg.Done()
		controllers.StartGRPCServer()
	}()

	fmt.Println("Starting servers...")
	wg.Wait()
}
