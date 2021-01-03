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
func Stage1(player player.Player) {
	stage1Map, len := Lev1Map()
	player.InitPlayer() // set player`s location (1,1) and hp 3.

	display.Map(stage1Map, player)

	for {
		move, _, _ := keyboard.GetSingleKey()
		player.Move(move, stage1Map)
		display.Map(stage1Map, player)
		px, py := player.Location()
		if end(px, py, len-2, len-2) {
			break
		}
	}

}
