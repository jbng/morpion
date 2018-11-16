package server

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"encoding/json"
)

var wsupgrader = websocket.Upgrader{}

type payload struct{
	event string
	tile_index int

}
func wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
       return
    }

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        var payload *payload
        if err := json.Unmarshal(msg, &payload); err != nil {
        	log.Fatal(err)
		}
		if (payload.event == "play_at_coordinate") {
			log.Println("play_at_coordinate: %d", payload.tile_index)
		}

		if (payload.event == "play_again") {
			log.Println("play_again")
		}

		conn.WriteMessage(t, msg)
    }
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	return r
}