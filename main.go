package main

import (
	"fmt"

	"github.com/eiannone/keyboard"

	"github.com/j1mmyson/maze_golang/player"
	"github.com/j1mmyson/maze_golang/stage"
)

func main() {
	player := player.Player{}
	player.InitPlayer()

	char, _, _ := keyboard.GetSingleKey()
	if char == 'w' {
		fmt.Println("w pressed")
	}
	fmt.Println(player)
	stage.Stage1(player)
}
