import React, { useEffect, useState } from 'react';
import Sketch from 'react-p5';
import p5Types from 'p5'; // Import this for type hinting

interface GameAreaProps {
  players: [
    {
      id: number;
      x: number;
      y: number;
    }
  ];
  playerId: number;
}

export default function GameArea({ players, playerId }: GameAreaProps) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  const maxSpeed = 3;

  // Update position whenever xPos or yPos changes
  useEffect(() => {
    setPosition({ x: 0, y: 0 });
  }, [players]);

  // const setup = (p5: p5Types, canvasParentRef: Element) => {
  //   p5.createCanvas(800, 600).parent(canvasParentRef); // Set a fixed size or adjust as needed
  // };

  // const draw = (p5: p5Types) => {
  //   p5.background(index === 0 ? 200 : 100);

  //    // let vel = index == 0 ? p5.createVector(p5.mouseX, p5.mouseY) : p5.createVector(x, y);
  //   // vel.sub(p5.createVector(position.x, position.y))
  //   // vel.setMag(maxSpeed)
  //   // setPosition({x: position.x + vel.x, y: position.y + vel.y})

    
  //   p5.fill(0);
  //   p5.rect(position.x, position.y, 50, 50); // Adjust size as needed
  // };

  // return <Sketch setup={setup} draw={draw} />;
  const setup = (p5: p5Types, canvasParentRef: Element) => {
    p5.createCanvas(p5.windowWidth, p5.windowHeight).parent(canvasParentRef);
  };

  const draw = (p5: p5Types) => {
    p5.background(200); // You can change this for a different background
    
    players.forEach((player, index) => {
      console.log(p5.mouseX, p5.mouseY)
      let vel = index == 0 ? p5.createVector(p5.mouseX, p5.mouseY) : p5.createVector(player.x, player.y);
      vel.sub(p5.createVector(position.x, position.y))
      vel.setMag(maxSpeed)
      setPosition({x: position.x + vel.x, y: position.y + vel.y})
      p5.fill(index === 0 ? 125 : 0); // Change color based on player
      p5.rect(index == playerId ? p5.mouseX + vel.x : 0, index == playerId ? p5.mouseY + vel.y : 0, 50, 50); // Draw the spaceship
    });
  };

  return <Sketch setup={setup} draw={draw} />;
}
