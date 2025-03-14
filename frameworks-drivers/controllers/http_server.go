package controllers

import (
	interface_adapters "go-boilerplate-clean-architecture/interface-adapters"
	"log"
	"net/http"
	"strconv"
)

func setupUserRoutes(ua *interface_adapters.UserAdapter) {
	http.HandleFunc("/users", ua.CreateUserAdapter)
}

func setupHelloWorldRoutes(hwa *interface_adapters.HelloWorldAdapter) {
	http.HandleFunc("/hello-world", hwa.HelloWorldAdapter)
}

type ServerHTTP struct {
	Port int `json:"port"`
}

func (o *ServerHTTP) ServeHTTP() {
	log.Printf("HTTP listening on port: %s", strconv.Itoa(o.Port))
	if err := http.ListenAndServe(":"+strconv.Itoa(o.Port), nil); err != nil {
		log.Fatalf("HTTP error starting server: %s", err)
	}

}

func NewServerHTTP(port int,
	hwController *interface_adapters.HelloWorldAdapter,
	uController *interface_adapters.UserAdapter) *ServerHTTP {
	r := &ServerHTTP{Port: port}
	setupHelloWorldRoutes(hwController)
	setupUserRoutes(uController)
	return r
}
