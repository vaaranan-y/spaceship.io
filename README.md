# spaceship.io

A multiplayer game being developed in Golang.

# Getting Started

- In one terminal go to the backend directory
  `cd ./backend`
- In this terminal, start up the backend server
  `go run *.go`
- In another, or multiple terminals, connect to the server via wscat
  `wscat -c ws://localhost:8080/join`
- Should receive a message on the server end for every terminal/player that joins the game
