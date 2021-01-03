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

// Move moves player as key-in
func (p *Player) Move(move rune, maze [][]byte) {

	maze[p.x][p.y] = 0

	switch move {
	case 'w':
		if maze[p.x-1][p.y] != '#' {
			p.x = p.x - 1
		}
	case 'a':
		if maze[p.x][p.y-1] != '#' {
			p.y = p.y - 1
		}
	case 'd':
		if maze[p.x][p.y+1] != '#' {
			p.y = p.y + 1
		}
	case 's':
		if maze[p.x+1][p.y] != '#' {
			p.x = p.x + 1
		}
	}
}

// Location return location of player
func (p *Player) Location() (int, int) {
	return p.x, p.y
}

// InitPlayer init player`s location.
func (p *Player) InitPlayer() {
	p.x = 1
	p.y = 1
	p.hp = 3
}
