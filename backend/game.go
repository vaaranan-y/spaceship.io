package backend

type Game struct {
	Players map[int64]*Player
}

func (Game g) GetPlayers() int64 {
	return g.Players
}

func CreateGame() *Game {
	return &Game{
		Players: make(map[int64]Player)
	}
}