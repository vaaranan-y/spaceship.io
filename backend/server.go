package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool { return true }, // Accept all clients for now
}

var game = CreateGame()
var newPlayerId int64
var playerCount int64


func playerConnectionEndpoint(w http.ResponseWriter, r *http.Request){
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
	newPlayerId += 1
	playerCount += 1
	newPlayerAlertMessage := fmt.Sprintf("A new player has joined!")
	for _, player := range game.Players {
		playerConn := player.Conn
		playerConn.WriteMessage(websocket.TextMessage, []byte(newPlayerAlertMessage))
	}

	game.AddPlayer(player)

	fmt.Printf("Player %v has joined!\n", player.ID)

	for {
		// Receive Message
		messageType, p, err := wsConn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Message Received: %s\n", p)
		messageContent := string(p)

		if(messageContent == "info"){
			// Echo Message
			message := fmt.Sprintf("There are currently %d player(s)", playerCount)
			err = wsConn.WriteMessage(messageType, []byte(message))
			if(err != nil){
				log.Fatal(err)
				return
			}
		}
	}

	defer wsConn.Close()
}

func gameLoop(game *Game) {
	tick := time.Tick(16*time.Millisecond)
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
		}
	}
}

func startServer(){
	http.HandleFunc("/join", playerConnectionEndpoint)

	// Start Server
	newPlayerId = 0
	playerCount = 0

	

	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}

func main(){
	// Set up routes
	go startServer()
	go gameLoop(game)
	select {}
}	