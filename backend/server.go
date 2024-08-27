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


func reader(conn *websocket.Conn){
	for {
		// Receive Message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Message Received: %s\n", p)

		// Echo Message
		err = conn.WriteMessage(messageType, p)
		if(err != nil){
			log.Fatal(err)
			return
		}
	}
}

func playerConnectionEndpoint(w http.ResponseWriter, r *http.Request){
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if(err != nil){
		http.Error(w, "Unsupported Method", http.StatusNotFound)
	} 
	
	player := &Player{
		ID: newPlayerId,
		NickName: "",
		GameID: 0,
		PosX: 0,
		PosY: 0,
		Health: 100,
		Damage: 5,
		Alive: true,
		Conn: wsConn,
	}
	newPlayerId = newPlayerId + 1
	
	game.AddPlayer(player)

	fmt.Printf("Player %v has joined!\n", player.ID)
	reader(wsConn)
	defer wsConn.Close()
}

func main(){
	// Set up routes
	http.HandleFunc("/join", playerConnectionEndpoint)

	// Start Server
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}