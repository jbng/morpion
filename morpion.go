package main

import (
	"morpion/server"
	"github.com/gin-gonic/contrib/static"
)
func main() {
	game := server.NewGame()
	router := server.InitRouter(game)

	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	// Start and run the server
	router.Run(":3000")
}
