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
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/join");

    ws.current.onmessage = (event) => {
      const data =  JSON.parse(event.data);
      switch(data.type) {
        case "positions": {
          setPlayers(Object.values(data.message))
          break
        }
        case "player_data": {
          setPlayerId(Number(data.message))
          console.log("VAARANAN: ", Number(data.message))
          break;
        }
      }
    }
    
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
        <GameArea players={players} playerId={playerId}/>
      }
    </div>
  );
}
