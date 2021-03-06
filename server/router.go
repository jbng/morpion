package server

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"encoding/json"
)

var wsupgrader = websocket.Upgrader{}

type Payload struct {
	Event      string `json:"event"`
	Tile_index string `json:"tile_index"`
}

func wshandler(g *Game, w http.ResponseWriter, r *http.Request) {

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	log.Println("websocket open")
	if g.IsSpotAvailable() {
		g.AddPlayer(conn)
		log.Println("Player", g.players_list)
	}
	for {
		t, msg, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		payload := new(Payload)
		if err := json.Unmarshal(msg, &payload); err != nil {
			log.Fatal(err)
		}
		log.Println("Message", string(msg))
		log.Println("Payload: ", payload.Event, " & ", payload.Tile_index)
		if (payload.Event == "play_at_coordinate") {
			log.Printf("play_at_coordinate: %s", payload.Tile_index)
		}

		if (payload.Event == "play_again") {
			log.Println("play_again")
		}

		conn.WriteMessage(t, msg)
	}
}

func InitRouter(g *Game) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(g, c.Writer, c.Request)
	})

	return r
}
