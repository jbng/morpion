package main

import (
	"morpion/server"
	"github.com/gin-gonic/contrib/static"
)
func main() {
	router := server.InitRouter()

	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	// Start and run the server
	router.Run(":3000")
}
