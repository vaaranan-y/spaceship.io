import React, { useState } from 'react';
import "./spaceship.css"

interface SpaceshipProps {
    x: number;
    y: number;
}

function Spaceship({x, y} : SpaceshipProps) {
    return (
        <div className="spaceship" style={{ left: x, top: y }}>
        </div>
    );
}

export default Spaceship;