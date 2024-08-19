package backend

type Game struct {
	Players []int64
}

func (Game g) getPlayers() int64 {
	return g.Players
}