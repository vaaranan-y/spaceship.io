package models

type Game struct {
	Players map[int64]*Player
	GameState string
}

func (g *Game) GetPlayers() map[int64]*Player {
	return g.Players
}

func CreateGame() *Game {
	return &Game{
		Players: make(map[int64]*Player),
		GameState: "waiting",
	}
}

func (g *Game) AddPlayer(playerToAdd *Player) bool{
	g.Players[playerToAdd.ID] = playerToAdd
	return true
}

func (g *Game) RemovePlayer(playerID int64) bool{
	_, exists := g.Players[playerID]
	if(exists){
		delete(g.Players, playerID)
		return true
	} else {
		return false
	}
	
}