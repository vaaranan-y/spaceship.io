import React, { useEffect, useRef, useState } from 'react';
import GameArea from './components/gamearea';
import './App.css';

interface coordinates {
  x: number;
  y: number;
}

interface Player {
  ID: number;
  NickName: string;
	GameID: number;
	PosX: number;
	PosY: number;
	Health: number;
	Damage: number;
	Alive: number;
	Color: string;
	Conn: any
}

export default function App() {
  const [players, setPlayers]  = useState<any>([])
  const [playerId, setPlayerId] = useState<number>(0);
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/join");
    // 
    ws.current.onmessage = (event) => {
      const data =  JSON.parse(event.data);
      switch(data.type) {
        case "update": {
          // console.log("POSITIONS: ", data);
          setPlayers(Object.values(data.players));
          break
        }
        // case "take_hit": {
        //   console.log("TAKE HIT: ", data);
        //   let oldColor = colors[playerId];
        //   colors[playerId] = "red";
        //   setTimeout(() => {
        //     colors[playerId] = oldColor;
        //   }, 200);
          
        //   break;
        // }
        case "init": {
          console.log("MESSAGE: ", data.playerId);
          setPlayerId(data.playerId);
        }

      }
    }

    // // I want to send the player's position to the server
    // const sendPlayerPosition = (position: coordinates) => {
    //   if (ws.current && ws.current.readyState === WebSocket.OPEN) {
    //   ws.current.send(JSON.stringify({ type: 'player_position', position }));
    //   }
    // };

    // // Example usage: sending a dummy position
    // sendPlayerPosition({ x: 100, y: 200 });
    
    ws.current.onopen = () => {
      console.log('WebSocket Connection Opened')
    }

    ws.current.onerror = () => {
      console.log('WebSocket Connection Error!')
    }


  }, [])

  return (
    <div className="App">
      {
        <GameArea playerId={playerId} players={players} ws={ws}/>
      }
    </div>
  );
}
