package server

import (
	"github.com/rs/xid"
	"github.com/gorilla/websocket"
)


type Game struct {
	players_list []xid.ID
	player_id_turn xid.ID
	has_winner bool
	board [9]int
}

type Player struct {
	id xid.ID
	cross bool
	conn *websocket.Conn
}

func NewGame() *Game {
	g := new(Game)
	g.players_list = make([]xid.ID, 2)
	g.has_winner = false
	g.board = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	return g
}

func (g *Game) resetGame() {
	g.has_winner = false
	g.board = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (g *Game) IsSpotAvailable() bool {
	return len(g.players_list) < 2
}

func (g *Game) AddPlayer(c *websocket.Conn) {
	p := new(Player)
	p.id = xid.New()
	g.players_list = append(g.players_list, p.id)
}