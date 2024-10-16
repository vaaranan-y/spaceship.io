import React, { MutableRefObject, useEffect, useState, useRef } from 'react';
import Sketch from 'react-p5';
import p5Types from 'p5';

interface GameAreaProps {
  players: Array<{
    id: number;
    x: number;
    y: number;
  }>;
  playerId: number;
  colors: Array<string>;
  ws: MutableRefObject<WebSocket | null>;
}

export default function GameArea({ players, playerId, colors, ws }: GameAreaProps) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  const maxSpeed = 0.01; // Lower value for smooth movement
  const positionRef = useRef(position); // Use ref to keep track of the latest position

  useEffect(() => {
    positionRef.current = position; // Update the ref whenever position changes
  }, [position]);

  const setup = (p5: p5Types, canvasParentRef: Element) => {
    p5.createCanvas(p5.windowWidth, p5.windowHeight).parent(canvasParentRef);
  };

  const draw = (p5: p5Types) => {
    p5.background(200);
    
    players.forEach((player, index) => {
      p5.fill(colors[index]);

      if (index === playerId) {
        const targetX = p5.mouseX;
        const targetY = p5.mouseY;

        // Interpolate the blob's position towards the mouse position
        const newX = positionRef.current.x + (targetX - positionRef.current.x) * maxSpeed;
        const newY = positionRef.current.y + (targetY - positionRef.current.y) * maxSpeed;
        setPosition({ x: newX, y: newY });

        p5.rect(newX, newY, 50, 50);

        // Send updated position to the WebSocket
        if (ws.current) {
          ws.current.send(JSON.stringify({ type: 'update_position', message: { id: playerId, x: newX, y: newY } }));
        }
      } else {
        // Draw other players at their positions
        p5.rect(player.x, player.y, 50, 50);
      }
    });
  };

  return <Sketch setup={setup} draw={draw} />;
}
