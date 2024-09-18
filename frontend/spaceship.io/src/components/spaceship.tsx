import React, { useState } from 'react';
import Sketch from 'react-p5';
import p5Types from 'p5'; // Import this for type hinting

export default function Soaceship() {
  const [position, setPosition] = useState({ x: 300, y: 300 });
  const maxSpeed = 3;
  
  const setup = (p5: p5Types, canvasParentRef: Element) => {
    p5.createCanvas(p5.windowWidth, p5.windowHeight).parent(canvasParentRef);
  };

  const draw = (p5: p5Types) => {
    p5.background(100);
    
    const dirX = p5.mouseX - position.x;
    const dirY = p5.mouseY - position.y;
    const angle = Math.atan2(dirY, dirX);
    
    const velX = Math.cos(angle) * maxSpeed;
    const velY = Math.sin(angle) * maxSpeed;

    setPosition(prev => ({
      x: prev.x + velX,
      y: prev.y + velY,
    }));

    p5.fill(0);
    p5.rect(position.x, position.y, 50, 50);
  };

  return <Sketch setup={setup} draw={draw} />;
};


