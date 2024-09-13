import React, { useEffect, useState } from 'react';
import Spaceship from './components/spaceship';
import './App.css';

function App() {
  const [x, setX] = useState<number>(0);
  const [y, setY] = useState<number>(0);


  useEffect(() => {
    const handleMouseMove = (event: MouseEvent) => {
      console.log(`{x: ${event.clientX}, y:${event.clientX}}`)
      setX(event.clientX)
      setY(event.clientY)
    }
    window.addEventListener('mousemove', handleMouseMove);
  }, [])


  
  return (
    <div className="App">
      <div className="TitleCard">
      SPACESHIP.IO
      </div>
      <Spaceship x={x} y={y}/>
    </div>
  );
}

export default App;
