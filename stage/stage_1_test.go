package stage_test

import (
	"testing"

	"github.com/j1mmyson/maze_golang/display"
	"github.com/j1mmyson/maze_golang/player"
	"github.com/j1mmyson/maze_golang/stage"
)

// Stage1 : escape
func TestStage1(t *testing.T) {
	stage1Map, _ := stage.Lev1Map()
	p := player.Player{}
	p.InitPlayer()
	display.Map(stage1Map, p)
}
