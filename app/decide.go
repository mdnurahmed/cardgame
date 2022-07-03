package app

import "fmt"

func (g *Game) Decide() (int, string, int) {
	winner := 0
	resultCard, resultFrequency := g.evaluate(0)
	fmt.Printf("Player 1's \n most frequent card has frequency %d\n most valuable card is %s\n", resultFrequency, resultCard)
	for i := 1; i < 4; i++ {
		// frequency of the most frequent card
		card, frequency := g.evaluate(i)
		fmt.Printf("Player %d's \n most frequent card has frequency %d\n most valuable card is %s\n", i+1, frequency, card)
		if frequency > resultFrequency || (frequency == resultFrequency && g.Rank[card] > g.Rank[resultCard]) {
			resultCard = card
			resultFrequency = frequency
			winner = i
		}
	}
	return winner + 1, resultCard, resultFrequency
}

func (g *Game) evaluate(p int) (string, int) {
	resCard, frequency := "", 0
	hashmap := make(map[string]int)
	max := make(map[string]string) // max rank of this type of card this player has
	for i := 0; i < len(g.Players[p]); i++ {
		card := g.Players[p][i]
		alpha := getAlphaNumericValue(card)
		hashmap[alpha]++
		if g.Rank[max[alpha]] < g.Rank[card] {
			max[alpha] = card
		}
		if hashmap[alpha] > frequency ||
			(frequency == hashmap[alpha] && g.Rank[max[alpha]] > g.Rank[resCard]) {
			resCard = max[alpha]
			frequency = hashmap[alpha]
		}
	}
	return resCard, frequency
}

func getAlphaNumericValue(card string) string {
	chars := []rune(card)
	return string(chars[0 : len(chars)-1])
}