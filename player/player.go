package player

// Player has location and hp of player.
type Player struct {
	x, y int
	hp   int
}

// User has name and score of user.
type User struct {
	name  string
	score int
}

// InitPlayer init player`s location.
func (p *Player) InitPlayer() {
	p.x = 1
	p.y = 1
	p.hp = 3
}
