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

interface Bullet {
  x: number;
  y: number;
  dx: number;
  dy: number;
}

export default function GameArea({ players, playerId, colors, ws }: GameAreaProps) {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  const [bullets, setBullets] = useState<Bullet[]>([]);
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

    // Move and draw bullets
    setBullets((prevBullets) => {
      let currBullets = prevBullets.map((bullet) => {
        const newX = bullet.x + bullet.dx;
        const newY = bullet.y + bullet.dy;

        // Check for collisions with other players
        let hasHit = false;
        players.forEach((player, index) => {
          if (index !== playerId && newX >= player.x && newX <= player.x + 50 && newY >= player.y && newY <= player.y + 50) {
            // Send collision information to the WebSocket
            if (ws.current) {
              ws.current.send(JSON.stringify({ type: 'bullet_hit', message: { shooterId: playerId, targetId: index } }));
            }
            hasHit = true;
          }
        });

        if (hasHit) {
          return null; // Mark bullet for removal
        } else {
          return { ...bullet, x: newX, y: newY };
        }
      }).filter((bullet): bullet is Bullet => bullet !== null); // Filter out null values

      return currBullets.filter((bullet) => {
        // Remove bullets that are out of bounds or have hit a player
        if (bullet.x < 0 || bullet.x > p5.width || bullet.y < 0 || bullet.y > p5.height) {
          return false;
        }
        return true;
      });
    });

    bullets.forEach((bullet) => {
      if (bullet) {
        p5.ellipse(bullet.x, bullet.y, 10, 10);
      }
    });
  };

  const mousePressed = (p5: p5Types) => {
    const angle = Math.atan2(p5.mouseY - positionRef.current.y, p5.mouseX - positionRef.current.x);
    const speed = 5;
    const dx = Math.cos(angle) * speed;
    const dy = Math.sin(angle) * speed;

    setBullets((prevBullets) => [...prevBullets, { x: positionRef.current.x, y: positionRef.current.y, dx, dy }]);
  };

  return <Sketch setup={setup} draw={draw} mousePressed={mousePressed} />;
}
