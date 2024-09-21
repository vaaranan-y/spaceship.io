import React, { useState } from 'react';
import Sketch from 'react-p5';
import p5Types from 'p5'; // Import this for type hinting
interface SpaceshipProps {
  index: number
  x: number
  y: number
}

export default function Soaceship({index, x, y} : SpaceshipProps) {
  const [position, setPosition] = useState({ x: 300, y: 300 });
  const maxSpeed = 3;
  
  const setup = (p5: p5Types, canvasParentRef: Element) => {
    p5.createCanvas(p5.windowWidth, p5.windowHeight).parent(canvasParentRef);
  };

  const draw = (p5: p5Types) => {
    p5.background(index == 0 ? 200 : 100);


    let vel = index == 0 ? p5.createVector(p5.mouseX, p5.mouseY) : p5.createVector(x, y);
    vel.sub(p5.createVector(position.x, position.y))
    vel.setMag(maxSpeed)
    setPosition({x: position.x + vel.x, y: position.y + vel.y})

    p5.fill(0);
    p5.rect(position.x, position.y, 50, 50);
  };

  return <Sketch setup={setup} draw={draw} />;
};


