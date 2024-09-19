import React, { useEffect, useState } from 'react';
import Spaceship from './components/spaceship';
import './App.css';

export default function App() {
  const [players, setPlayers]  = useState([])
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
