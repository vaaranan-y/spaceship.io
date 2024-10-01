package api
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"spaceshipio/backend/internal/game"
	"spaceshipio/backend/internal/models"
	"encoding/json"
)

type Message struct {
    Type string `json:"type"` // Use struct tags to match JSON keys
    Message string `json:"message"` // Example additional field
}

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

	// Create Player
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

	// Add player to game Manager
	gameManager.Mu.Lock()
	gameManager.Game.AddPlayer(newPlayer)
	gameManager.Players = append(gameManager.Players, newPlayer)
	gameManager.PlayerCount++
	gameManager.Mu.Unlock()
	fmt.Printf("Player %v has joined!\n", newPlayer.ID)
	fmt.Printf("There are currently %v players in the game\n", gameManager.PlayerCount)

	
	if(gameManager.PlayerCount >= 3){
		newPlayerAlertMessage := jsonifyData(map[string]string{
			"type":    "player_alert",
			"message": "At least three players have joined, the game will begin momentarily!",
		})
	
		for _, player := range gameManager.Players {
			playerConn := player.Conn
			playerConn.WriteMessage(websocket.TextMessage, newPlayerAlertMessage)
		}
	}
	go handleMessages(newPlayer, gameManager)
}

func handleMessages(player *models.Player, gameManager *gameManager.GameManager) {
	defer player.Conn.Close() 

	for {
		messageType, msg, err := player.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message from player %v: %v\n", player.ID, err)
			break
		}

		// Process the received message
		fmt.Printf("Received message from player %v\n", player.ID)
		var unmarshalledMsg Message
		json.Unmarshal(msg, &unmarshalledMsg)
		receivedType := unmarshalledMsg.Type

		gameManager.Mu.Lock()
		switch receivedType {
		case "info":
			message := jsonifyData(map[string]string{
				"type":    "info",
				"message": fmt.Sprintf("There are currently %d player(s)", len(gameManager.Players)),
			})
			err = player.Conn.WriteMessage(messageType, message)
			if(err != nil){
				log.Fatal(err)
				return
			}
		case "positions":
			positions := make(map[string]map[string]float64)
			for _, p := range gameManager.Players {
				positions[fmt.Sprintf("Player%d", p.ID)] = map[string]float64{
					"x": p.PosX,
					"y": p.PosY,
				}
			}
			message := jsonifyData(map[string]interface{}{
				"type":    "positions",
				"message": positions,
			})
			
			err = player.Conn.WriteMessage(messageType, message)
			if(err != nil){
				log.Fatal(err)
				return
			}
		default:
			message := jsonifyData(map[string]string{
				"type":    "unknown",
				"message": fmt.Sprintf("Message unknown"),
			})

			err = player.Conn.WriteMessage(messageType, message)
			if(err != nil){
				log.Fatal(err)
				return
			}
			// var posnMsg PositionMessage
			// err = json.Unmarshal(p, &posnMsg)
			// if err != nil {
			// 	log.Println("Error: ", err)
			// 	continue
			// }
			// if posnMsg.Type == "update_position" {
			// 	game.Players[posnMsg.ID].PosX = posnMsg.X
			// 	game.Players[posnMsg.ID].PosY = posnMsg.Y
			// }

			// for _, player := range game.Players {
			// 	playerConn := player.Conn
			// 	playerConn.WriteMessage(websocket.TextMessage, jsonifyData(posnMsg))
			// }
	}
		gameManager.Mu.Unlock()
	}
}
