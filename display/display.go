package display

import (
	"fmt"

	"github.com/j1mmyson/maze_golang/player"
)

// Start display start screen of game.
func Start() {

}

// Map display map.
func Map(m [][]byte, p player.Player) {
	length := len(m)
	px, py := p.Location()
	m[px][py] = 'P'
	fmt.Print("\033[H\033[2J")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if m[i][j] == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%2c", m[i][j])
			}

		}
		fmt.Println()
	}
}
