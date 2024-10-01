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
      const data =  JSON.parse(event.data);
      switch(data.type) {
        case "positions": {
          setPlayers(Object.values(data.message))
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
          return (
            <Spaceship index={index} xPos={index*100} yPos={index*100}/>
          );
        })
      }
    </div>
  );
}
