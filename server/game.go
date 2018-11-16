package server

type Game struct {
	board [9]int
}

func NewGame() *Game {
	game := new(Game)
	game.board = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	return game
}
