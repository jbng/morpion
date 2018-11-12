package main

import (
	"morpion/router"
	"github.com/gin-gonic/contrib/static"
)
func main() {
	router := router.InitRouter()

	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	// Start and run the server
	router.Run(":3000")
}
