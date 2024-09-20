import React, { useEffect, useRef, useState } from 'react';
import Spaceship from './components/spaceship';
import './App.css';

export default function App() {
  const [players, setPlayers]  = useState([])
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/join");
    ws.current.onmessage = (event) => {
      const message = event.data;
      console.log(message)
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
        players.map((player, index) => {
          return (
            <Spaceship/>
          );
        })
      }
    </div>
  );
}
