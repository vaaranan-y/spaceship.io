package main

import (
	"fmt"
	"log"
	"net/http"
	"spaceshipio/backend/internal/game"
	"spaceshipio/backend/api"
)

func setUpRoutes(gameManager *gameManager.GameManager){
	// Add all routes here
	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
        api.PlayerConnectionEndpoint(gameManager, w, r)
    })
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
	gameManager := gameManager.CreateGameManager()
	setUpRoutes(gameManager)
	go startServer()
	select {}
}	