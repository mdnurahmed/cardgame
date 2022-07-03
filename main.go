package main

import (
	"cardgame/app"
	"fmt"
)

func main() {
	game := app.NewGame()
	game.Play()
	winner, _, _ := game.Decide()
	fmt.Printf("\nwinner is player %d\n", winner)
}
