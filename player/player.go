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

// S1Move : movement of stage1
func (p *Player) S1Move(move rune, maze [][]byte) {

	maze[p.x][p.y] = 0

	switch move {
	case 'w':
		for {
			if maze[p.x-1][p.y] == '#' {
				break
			}
			p.x = p.x - 1
		}
	case 'a':
		for {
			if maze[p.x][p.y-1] == '#' {
				break
			}
			p.y = p.y - 1
		}
	case 'd':
		for {
			if maze[p.x][p.y+1] == '#' {
				break
			}
			p.y = p.y + 1
		}
	case 's':
		for {
			if maze[p.x+1][p.y] == '#' {
				break
			}
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
