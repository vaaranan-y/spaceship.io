package models

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	ID int64
	NickName string
	GameID int64
	PosX float64
	PosY float64
	Health int64
	Damage int64
	Alive bool
	Color string
	Conn *websocket.Conn
}

func (p *Player) getHealth() int64 {
	return p.Health
}

func (p *Player) getDamageCapability() int64 {
	return p.Damage
}

func (p *Player) getCoordinates() (float64, float64) {
	return p.PosX, p.PosY
}

func (p *Player) takeDamage(damage int64) {
	if(damage > p.Health){
		p.Alive = false
	} else {
		p.Health -= damage
	}
}