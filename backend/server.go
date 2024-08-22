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
}

func reader(connection *websocket.Conn){
	for {
		messageType, p, err := connection.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(string(p)) // Print out message
		err = connection.WriteMessage(messageType, p)
		if(err != nil){
			log.Fatal(err)
			return
		}
	}
}

func testEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func (r *http.Request) bool { return true } // Accept all clients for now

	if(r.URL.Path != "/hello"){
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if(r.Method != "GET"){
		http.Error(w, "Unsupported Method", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, world!")

	ws, err := upgrader.Upgrade(w, r, nil)
	if(err != nil){
		http.Error(w, "Unsupported Method", http.StatusNotFound)
	} 
	reader(ws)
}

func main(){
	// Set up routes
	http.HandleFunc("/hello", testEndpoint)

	// Start Server
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}