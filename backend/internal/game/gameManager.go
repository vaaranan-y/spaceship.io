package gameManager

import (
	"log"
	"time"
	"sync"
	"spaceshipio/backend/internal/models"
)

type GameManager struct {
	Game *models.Game
	Players []*models.Player  // Track connected players
	PlayerCount int64
	Mu      sync.Mutex        // Protect the player map for concurrency
}

func CreateGameManager() *GameManager {
	manager := &GameManager{
		Game: models.CreateGame(),
		Players: make([]*models.Player, 0),
		PlayerCount: 0,
	}
	// Start the game loop as a goroutine
	go manager.StartGameLoop()
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


func (gm *GameManager) StartGameLoop() {
	ticker := time.Tick(1000*time.Millisecond) // 20 FPS (adjust as necessary)

	for t := range ticker {
		log.Printf("Hello world! %v\n", t)
		// gm.updateGameState()
	}
}