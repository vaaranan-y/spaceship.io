package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
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

	// err = wsConn.WriteMessage(messageType, []byte(message))
	// if(err != nil){
	// 	log.Fatal(err)
	// 	return
	// }
	
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

func main(){
	// Set up routes
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