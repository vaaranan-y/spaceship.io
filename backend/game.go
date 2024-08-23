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

func (Game g) AddPlayer(playerToAdd *Player) bool{
	g.Players[playerToAdd.ID] = playerToAdd
	return true
}

func (Game g) RemovePlayer(playerID int64) bool{
	player, exists := g.Players[playerID]
	if(exists){
		delete(g.Players, playerID)
	} else {
		return false
	}
	
}