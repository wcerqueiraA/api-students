package main

import (
	"github.com/wcerqueiraA/api-students/api"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	server.Start()
}
