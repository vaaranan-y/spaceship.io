import React, { useEffect, useRef, useState } from 'react';
import GameArea from './components/gamearea';
import './App.css';

interface coordinates {
  x: number;
  y: number;
}

interface playerInfo {
  id: number;
  x: number;
  y: number;
}

export default function App() {
  const [playerId, setPlayerId] = useState<number>(0);
  const [players, setPlayers]  = useState<any>([])
  const [colors, setColors]  = useState<any>([])
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/join");

    ws.current.onmessage = (event) => {
      const data =  JSON.parse(event.data);
      switch(data.type) {
        case "positions": {
          // console.log("POSITIONS: ", data);
          setPlayers(Object.values(data.message));
          setColors(Object.values(data.colors));
          break
        }
        case "player_data": {
          setPlayerId(Number(data.message))
          console.log("VAARANAN: ", data);
          break;
        }
        case "take_hit": {
          console.log("TAKE HIT: ", data);
          let oldColor = colors[playerId];
          colors[playerId] = "red";
          setTimeout(() => {
            colors[playerId] = oldColor;
          }, 200);
          
          break;
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
        <GameArea players={players} playerId={playerId} colors={colors} ws={ws}/>
      }
    </div>
  );
}
