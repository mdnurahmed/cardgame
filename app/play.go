package app

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Game) Play() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.Cards), func(i, j int) { g.Cards[i], g.Cards[j] = g.Cards[j], g.Cards[i] })
	for i := 0; i < len(g.Cards); i += 4 {
		for j := 0; j < 4; j++ {
			g.Players[j] = append(g.Players[j], g.Cards[i+j])
		}
	}
	for i := 0; i < 4; i++ {
		fmt.Printf("player %d got the cards :\n", i+1)
		for j := 0; j < len(g.Players[i]); j++ {
			fmt.Println(g.Players[i][j])
		}
		fmt.Println()
	}
}