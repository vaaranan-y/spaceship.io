package backend

type Player struct {
	ID int64
	NickName string
	GameID int64
	PosX float64
	PosY float64
	Health int64
	Damage int64
	Alive bool
	Conn *websocket.Conn
}

func (Player p) getHealth() int64 {
	return p.Health
}

func (Player p) getDamageCapability() int64 {
	return p.Damage
}

func (Player p) getCoordinates() (int64, int64) {
	return PosX, PosY
}

func (Player *p) takeDamage(damage int64) int64 {
	if(damage > p.Health){
		p.Alive = false
	} else {
		p.Health -= damage
	}
}