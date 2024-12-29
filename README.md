# spaceship.io

## About

`spaceship.io` is a multiplayer game similar to [agar.io](https://agar.io/). In this game, each user gets their own ship, and can traverse the displayed map to shoot down other players playing the game. The stack for this game consists of the following:

- Backend
  - **Golang**: golang was used to create the backend server for handling and maintaining the server's start up, the game loop, and individual player data (using websockets). These tasks are all handled concurrently using separate `Goroutines`
- Frontend
  - **React.js**: this dynamic frontend framework, actively updates the UI from the perspective all the different users' movement, in realtime. This allows the user to perform their own actions with respect to the positions and actions of other players with accuracy and consistency amongst all users.
  - **p5.js**: this is an animation framework used to make the user experience more interactive (i.e. the user's avatar/ship will follow their mouse on the screen, similar to agar.io)
  - **TypeScript**: this was used over JavaScript for its numerous benefits, including static typing, readabilty, and code quality.

## Getting Started

### Backend

- In one terminal go to the backend directory
  `cd ./backend`
- In this terminal, start up the backend server
  `go run ./cmd/app`
- In another, or multiple terminals, connect to the server via wscat
  `wscat -c ws://localhost:8080/join`
- Should receive a message on the server end for every terminal/player that joins the game

### Frontend

- In another terminal, go to the frontend directory, under the spaceship.io project
  `cd ./frotnend/spaceship.io`
- In this terminal, start up the frontend server
  `npm start`
- The above will run the server on http://localhost:3000. Open this link in your preferred web browser

After setting up your backend and frontend servers, you should be able to begin testing and playing spaceship.io (locally)!
