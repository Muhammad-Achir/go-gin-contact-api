package main

import (
	"go-gin-contact-api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run()
}
