package stage

import (
	"github.com/j1mmyson/maze_golang/display"
	"github.com/j1mmyson/maze_golang/player"
)

// Boss : Boss stage
func Boss(p player.Player) {

	bossMap, _ := BossMap()
	p.InitPlayer()

	display.Map(bossMap, p)

	return
}
