package stage

import (
	"github.com/eiannone/keyboard"
	"github.com/j1mmyson/maze_golang/display"
	"github.com/j1mmyson/maze_golang/player"
)

func end(x, y, ex, ey int) bool {
	if x == ex && y == ey {
		return true
	}
	return false
}

// Stage1 : escape
func Stage1(p player.Player) {
	stage1Map, len := Lev1Map()
	p.InitPlayer() // set player`s location (1,1) and hp 3.

	display.Map(stage1Map, p)

	for {
		move, _, _ := keyboard.GetSingleKey()
		p.S1Move(move, stage1Map)
		display.Map(stage1Map, p)
		px, py := p.Location()
		if end(px, py, len-2, len-2) {
			break
		}
	}

}
