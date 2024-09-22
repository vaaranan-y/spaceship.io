package api
import (
	"fmt"
	// "log"
	"net/http"
	"github.com/gorilla/websocket"
	"spaceshipio/backend/internal/game"
	"spaceshipio/backend/internal/models"
	"encoding/json"
	// "time"
	// "encoding/json"

	// "spaceshipio/backend/internal/game"
	// "spaceshipio/backend/internal/api"
)

func jsonifyData(message interface{}) []byte {
	jsonMsg, err := json.Marshal(message)
    if err != nil {
		fmt.Printf("Error jsonifying message!");
        return nil
    }
    return jsonMsg
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool { return true }, // Accept all clients for now
}

func PlayerConnectionEndpoint(gameManager *gameManager.GameManager, w http.ResponseWriter, r *http.Request){
	// Create WebSocket Connection for Player
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if(err != nil){
		http.Error(w, "Unsupported Method", http.StatusNotFound)
	} 
	newPlayer := &models.Player{
		ID: gameManager.PlayerCount,
		NickName: "Temp",
		GameID: 0,
		PosX: 0,
		PosY: 0,
		Health: 100,
		Damage: 5,
		Alive: true,
		Conn: wsConn,
	}

	gameManager.PlayerCount += 1
	newPlayerAlertMessage := jsonifyData(map[string]string{
		"type":    "player_alert",
		"message": "At least three players have joined, the game will begin momentarily!",
	})

	for _, player := range gameManager.Players {
		playerConn := player.Conn
		playerConn.WriteMessage(websocket.TextMessage, newPlayerAlertMessage)
	}

	gameManager.Game.AddPlayer(newPlayer)
	fmt.Printf("Player %v has joined!\n", newPlayer.ID)

	// // Accept messages from the player
	// for {
	// 	// Receive Message
	// 	messageType, p, err := wsConn.ReadMessage()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	fmt.Printf("Message Received: %s\n", p)
	// 	messageContent := string(p)

	// 	// Handle Message
	// 	switch messageContent {
	// 		case "info":
				
	// 			message := jsonifyData(map[string]string{
	// 				"type":    "info",
	// 				"message": fmt.Sprintf("There are currently %d player(s)", playerCount),
	// 			})
	// 			err = wsConn.WriteMessage(messageType, message)
	// 			if(err != nil){
	// 				log.Fatal(err)
	// 				return
	// 			}
	// 		default:
	// 			var posnMsg PositionMessage
	// 			err = json.Unmarshal(p, &posnMsg)
	// 			if err != nil {
	// 				log.Println("Error: ", err)
	// 				continue
	// 			}
	// 			if posnMsg.Type == "update_position" {
	// 				game.Players[posnMsg.ID].PosX = posnMsg.X
	// 				game.Players[posnMsg.ID].PosY = posnMsg.Y
	// 			}

	// 			for _, player := range game.Players {
	// 				playerConn := player.Conn
	// 				playerConn.WriteMessage(websocket.TextMessage, jsonifyData(posnMsg))
	// 			}
	// 	}
	// }

	// defer wsConn.Close()
}