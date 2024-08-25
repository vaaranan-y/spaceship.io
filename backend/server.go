package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	// "./game"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool { return true }, // Accept all clients for now
}

// var game = game.CreateGame()

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

func webSocketEndpoint(w http.ResponseWriter, r *http.Request){
	wsConn, err := upgrader.Upgrade(w, r, nil)

	if(err != nil){
		http.Error(w, "Unsupported Method", http.StatusNotFound)
	} 

	reader(wsConn)
}

func testEndpoint(w http.ResponseWriter, r *http.Request) {

	if(r.URL.Path != "/hello"){
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if(r.Method != "GET"){
		http.Error(w, "Unsupported Method", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, world!")

}

func main(){
	// Set up routes
	http.HandleFunc("/hello", testEndpoint)
	http.HandleFunc("/ws", webSocketEndpoint)

	// Start Server
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}