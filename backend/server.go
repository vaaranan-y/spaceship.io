package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool { return true }, // Accept all clients for now
}
var game = CreateGame()
var newPlayerId int64
var playerCount int64
var gameStarted bool
var gameStartTime time.Time

type PositionMessage struct {
	Type string  `json:"type"`
	ID   int64 `json:"id"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

func jsonifyData(message interface{}) []byte {
	jsonMsg, err := json.Marshal(message)
    if err != nil {
		fmt.Printf("Error jsonifying message!");
        return nil
    }
    return jsonMsg
}


func playerConnectionEndpoint(w http.ResponseWriter, r *http.Request){
	// Create WebSocket Connection for Player
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if(err != nil){
		http.Error(w, "Unsupported Method", http.StatusNotFound)
	} 
	player := &Player{
		ID: newPlayerId,
		NickName: "Temp",
		GameID: 0,
		PosX: 0,
		PosY: 0,
		Health: 100,
		Damage: 5,
		Alive: true,
		Conn: wsConn,
	}

	// Update global data for application, and alert all players of new player
	newPlayerId += 1
	playerCount += 1
	newPlayerAlertMessage := jsonifyData(map[string]string{
		"type":    "player_alert",
		"message": "At least three players have joined, the game will begin momentarily!",
	})

	for _, player := range game.Players {
		playerConn := player.Conn
		playerConn.WriteMessage(websocket.TextMessage, newPlayerAlertMessage)
	}

	game.AddPlayer(player)
	fmt.Printf("Player %v has joined!\n", player.ID)

	// Accept messages from the player
	for {
		// Receive Message
		messageType, p, err := wsConn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Message Received: %s\n", p)
		messageContent := string(p)

		// Handle Message
		switch messageContent {
			case "info":
				
				message := jsonifyData(map[string]string{
					"type":    "info",
					"message": fmt.Sprintf("There are currently %d player(s)", playerCount),
				})
				err = wsConn.WriteMessage(messageType, message)
				if(err != nil){
					log.Fatal(err)
					return
				}
			default:
				var posnMsg PositionMessage
				err = json.Unmarshal(p, &posnMsg)
				if err != nil {
					log.Println("Error: ", err)
					continue
				}
				if posnMsg.Type == "update_position" {
					game.Players[posnMsg.ID].PosX = posnMsg.X
					game.Players[posnMsg.ID].PosY = posnMsg.Y
				}
		}
	}

	defer wsConn.Close()
}

func gameLoop(game *Game) {
	tick := time.Tick(1000*time.Millisecond)
	alertStart := false

	for t := range tick {
		fmt.Printf("Game Loop Update: %v\n", t)
		
		if(playerCount >= 3 && !alertStart){
			alertStart = true
			newPlayerAlertMessage := fmt.Sprintf("At least three players have joined, the game will begin momentarily!")
			for _, player := range game.Players {
				playerConn := player.Conn
				playerConn.WriteMessage(websocket.TextMessage, []byte(newPlayerAlertMessage))
			}
			gameStarted = true
			gameStartTime = time.Now().Add(5 * time.Second)
		} else if(gameStarted && time.Now().After(gameStartTime)) {
			playerPositionsMsg := "Players\n"
			for _, player := range game.Players {
				playerPositionsMsg += fmt.Sprintf("Player %v Coordinates: (%v, %v) ", player.ID, player.PosX, player.PosY)
			}
			for _, player := range game.Players {
				playerConn := player.Conn
				playerConn.WriteMessage(websocket.TextMessage, []byte(playerPositionsMsg))
			}

		}
	}
}

func setUpRoutes(){
	// Add all routes here
	http.HandleFunc("/join", playerConnectionEndpoint)
}

func startServer(){
	// Start the server on port 8080
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}

func main(){
	// Set up routes
	newPlayerId = 0
	playerCount = 0
	setUpRoutes()
	go startServer()
	go gameLoop(game)
	select {}
}	