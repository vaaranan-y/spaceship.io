import React, { useEffect, useState } from 'react';
import Sketch from 'react-p5';
import p5Types from 'p5'; // Import this for type hinting

interface SpaceshipProps {
  index: number;
  xPos: number;
  yPos: number;
}

export default function Spaceship({ index, xPos, yPos }: SpaceshipProps) {
  const [position, setPosition] = useState({ x: xPos, y: yPos });
  const maxSpeed = 3;

  // Update position whenever xPos or yPos changes
  useEffect(() => {
    setPosition({ x: xPos, y: yPos });
  }, [xPos, yPos]);

  const setup = (p5: p5Types, canvasParentRef: Element) => {
    p5.createCanvas(800, 600).parent(canvasParentRef); // Set a fixed size or adjust as needed
  };

  const draw = (p5: p5Types) => {
    p5.background(index === 0 ? 200 : 100);

     // let vel = index == 0 ? p5.createVector(p5.mouseX, p5.mouseY) : p5.createVector(x, y);
    // vel.sub(p5.createVector(position.x, position.y))
    // vel.setMag(maxSpeed)
    // setPosition({x: position.x + vel.x, y: position.y + vel.y})

    
    p5.fill(0);
    p5.rect(position.x, position.y, 50, 50); // Adjust size as needed
  };

  return <Sketch setup={setup} draw={draw} />;
}
