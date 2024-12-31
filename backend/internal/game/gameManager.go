package gameManager

import (
	// "log"
	// "time"
	"sync"
	"spaceshipio/backend/internal/models"
)

type GameManager struct {
	Type string `json:"type"` // Use struct tags to match JSON keys
	Players []*models.Player `json:"players"`
    Mu   sync.Mutex 
}

func CreateGameManager() *GameManager {
	manager := &GameManager{
		Type: "update",
		Players: make([]*models.Player, 0),
	}
	return manager
}

// func (gm *GameManager) RegisterPlayer(conn *websocket.Conn) *Player {
// 	gm.mu.Lock()
// 	defer gm.mu.Unlock()

// 	player := NewPlayer(conn)
// 	gm.players[player] = true
// 	log.Printf("Player joined: %v\n", player.ID)
// 	return player
// }

// // UnregisterPlayer removes a player from the game.
// func (gm *GameManager) UnregisterPlayer(player *Player) {
// 	gm.mu.Lock()
// 	defer gm.mu.Unlock()

// 	delete(gm.players, player)
// 	log.Printf("Player left: %v\n", player.ID)
// }
