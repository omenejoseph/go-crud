package main

import (
	"github.com/omenejoseph/go-crud/inititalizers"
	"github.com/omenejoseph/go-crud/routes"
)

func init() {
	inititalizers.LoadEnvVariables()
	inititalizers.ConnectToDB()
}

func main() {
	routes.HandleRouting()
}
