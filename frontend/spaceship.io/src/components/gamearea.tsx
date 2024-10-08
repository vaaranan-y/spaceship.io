import React, { MutableRefObject, useEffect, useState } from 'react';
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
  colors: [string];
  ws: MutableRefObject<WebSocket | null>;
}

export default function GameArea({ players, playerId, colors, ws }: GameAreaProps) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  const maxSpeed = 3;

  // return <Sketch setup={setup} draw={draw} />;
  const setup = (p5: p5Types, canvasParentRef: Element) => {
    p5.createCanvas(p5.windowWidth, p5.windowHeight).parent(canvasParentRef);
  };

  const draw = (p5: p5Types) => {
    p5.background(200); // You can change this for a different background
    
    players.forEach((player, index) => {
      // console.log(p5.mouseX, p5.mouseY)
      p5.fill(colors[index]);
      if(index == playerId) {
        p5.rect(p5.mouseX, p5.mouseY, 50, 50);
        if (ws.current) {
          ws.current.send(JSON.stringify({ type: 'update_position', message: { id: playerId, x: p5.mouseX, y: p5.mouseY } }));
        }
      } else {
        p5.rect(player.x, player.y, 50, 50);
      }
      // let vel = index == playerId ? p5.createVector(p5.mouseX, p5.mouseY) : p5.createVector(player.x, player.y);
      // vel.sub(p5.createVector(position.x, position.y))
      // vel.setMag(maxSpeed)
      // setPosition({x: position.x + vel.x, y: position.y + vel.y})
      // p5.fill(125); // Change color based on player
      // p5.rect(index == playerId ? p5.mouseX + vel.x : 0, index == playerId ? p5.mouseY + vel.y : 0, 50, 50); // Draw the spaceship
    });
  };

  return <Sketch setup={setup} draw={draw} />;
}
