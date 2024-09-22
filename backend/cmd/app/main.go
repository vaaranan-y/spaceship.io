package main

import (
	"fmt"
	"log"
	"net/http"
	"spaceshipio/backend/internal/game"
	"spaceshipio/backend/api"
)


// var gameStarted bool
// var gameStartTime time.Time

// type PositionMessage struct {
// 	Type string  `json:"type"`
// 	ID   int64 `json:"id"`
// 	X    float64 `json:"x"`
// 	Y    float64 `json:"y"`
// }

// type PlayerPositions struct {
// 	Type 	string					`json:"type"`
// 	Players map[int64](map[string](float64))	`json:"players"`
// }



// func gameLoop(game *Game) {
// 	tick := time.Tick(1000*time.Millisecond)
// 	alertStart := false

// 	for t := range tick {
// 		fmt.Printf("Game Loop Update: %v\n", t)
		
// 		if(playerCount >= 2 && !alertStart){
// 			alertStart = true
// 			gameStarted = true
// 			gameStartTime = time.Now().Add(5 * time.Second)
// 		} else if(gameStarted && time.Now().After(gameStartTime)) {

			
// 			playerMap := PlayerPositions{
// 				Type: "PlayerPositions",
// 				Players: make(map[int64](map[string](float64))),
// 			}
	

// 			for _, player := range game.Players {
// 				coordinates := map[string]float64{
// 					"x":player.PosX,
// 					"y":player.PosY,
// 				}
// 				playerMap.Players[player.ID] = coordinates
				
// 			}


// 			playerPositionsMsg := jsonifyData(playerMap)
// 			for _, player := range game.Players {
// 				playerConn := player.Conn
// 				playerConn.WriteMessage(websocket.TextMessage, playerPositionsMsg)
// 			}

// 		}
// 	}
// }

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
	// go gameLoop(game)
	select {}
}	