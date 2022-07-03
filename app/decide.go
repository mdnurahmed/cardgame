package app

func (g *Game) Decide() (int, string, int) {

	panic("not implemeneted")
}

func (g *Game) evaluate(p int) (string, int) {
	panic("not implemeneted")
}

func getAlphaNumericValue(card string) string {
	chars := []rune(card)
	return string(chars[0 : len(chars)-1])
}