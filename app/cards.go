package app

type Game struct {
	Cards   []string
	Players [][]string
	Rank    map[string]int
}

func NewGame() *Game {
	game := Game{}

	game.Cards = []string{"2@", "2#", "2^", "2*", "3@", "3#", "3^", "3*", "4@", "4#", "4^", "4*", "5@", "5#", "5^", "5*", "6@", "6#", "6^", "6*", "7@", "7#", "7^", "7*", "8@", "8#", "8^", "8*", "9@", "9#", "9^", "9*", "10@", "10#", "10^", "10*", "J@", "J#", "J^", "J*", "Q@", "Q#", "Q^", "Q*", "K@", "K#", "K^", "K*", "A@", "A#", "A^", "A*"}
	game.Rank = make(map[string]int)
	for i := 0; i < len(game.Cards); i++ {
		game.Rank[game.Cards[i]] = i + 1
	}
	game.Players = make([][]string, 4)
	for i := 0; i < 4; i++ {
		game.Players[i] = make([]string, 0)
	}
	return &game
}
