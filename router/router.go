package router

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"	
)

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
       return
    }

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
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