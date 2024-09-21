import React, { useEffect, useRef, useState } from 'react';
import Spaceship from './components/spaceship';
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
  const [players, setPlayers]  = useState<any>([])
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/join");
    ws.current.onmessage = (event) => {
      const message =  JSON.parse(event.data);
      switch(message.type) {
        case "PlayerPositions": {
          setPlayers(Object.values(message.players))
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
        players.map((player:playerInfo, index: number) => {
          console.log(index, player)
          return (
            <Spaceship index={index} x={player["x"]} y={player["y"]}/>
          );
        })
      }
    </div>
  );
}
