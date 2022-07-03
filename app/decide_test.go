package app

import (
	"testing"
)

func TestDecide1(t *testing.T){
	game := NewGame()
	game.Players[0] = []string{"K@","K#", "K*"}
	game.Players[1] = []string{"A@", "A#", "A^"}
	game.Players[2] = []string{}
	game.Players[3] = []string{}
    winner , resultCard, resultFrequency := game.Decide()
	expectedWinner , expectResultCard, expectedResultFrequency := 2,"A^",3
    if winner!=expectedWinner || resultCard !=expectResultCard || resultFrequency != expectedResultFrequency {
        t.Errorf("got (winner,resultCard,resultFrequency) = (%d,%s,%d), wanted(winner,resultCard,resultFrequency) = (%d,%s,%d)",winner,resultCard,resultFrequency,expectedWinner,expectResultCard,expectedResultFrequency )
    }
}

func TestDecide2(t *testing.T){
	game := NewGame()
	game.Players[0] = []string{"A@","A*"}
	game.Players[1] = []string{"A#", "A^"}
	game.Players[2] = []string{}
	game.Players[3] = []string{}
    winner , resultCard, resultFrequency := game.Decide()
	expectedWinner , expectResultCard, expectedResultFrequency := 1,"A*",2
    if winner!=expectedWinner || resultCard !=expectResultCard || resultFrequency != expectedResultFrequency {
        t.Errorf("got (winner,resultCard,resultFrequency) = (%d,%s,%d), wanted(winner,resultCard,resultFrequency) = (%d,%s,%d)",winner,resultCard,resultFrequency,expectedWinner,expectResultCard,expectedResultFrequency )
    }
}

func TestDecide3(t *testing.T){
	game := NewGame()
	game.Players[0] = []string{"A@","A*","2@","2#","2^"}
	game.Players[1] = []string{"A#", "A^"}
	game.Players[2] = []string{}
	game.Players[3] = []string{}
    winner , resultCard, resultFrequency := game.Decide()
	expectedWinner , expectResultCard, expectedResultFrequency := 1,"2^",3
    if winner!=expectedWinner || resultCard !=expectResultCard || resultFrequency != expectedResultFrequency {
        t.Errorf("got (winner,resultCard,resultFrequency) = (%d,%s,%d), wanted(winner,resultCard,resultFrequency) = (%d,%s,%d)",winner,resultCard,resultFrequency,expectedWinner,expectResultCard,expectedResultFrequency )
    }
}

func TestDecide4(t *testing.T){
	game := NewGame()
	game.Players[0] = []string{"2@","2#","A@","A*"}
	game.Players[1] = []string{"A#", "A^"}
	game.Players[2] = []string{}
	game.Players[3] = []string{}
    winner , resultCard, resultFrequency := game.Decide()
	expectedWinner , expectResultCard, expectedResultFrequency := 1,"A*",2
    if winner!=expectedWinner || resultCard !=expectResultCard || resultFrequency != expectedResultFrequency {
        t.Errorf("got (winner,resultCard,resultFrequency) = (%d,%s,%d), wanted(winner,resultCard,resultFrequency) = (%d,%s,%d)",winner,resultCard,resultFrequency,expectedWinner,expectResultCard,expectedResultFrequency )
    }
}


func TestGetAlphaNumericValue1(t *testing.T){
	alpha := getAlphaNumericValue("A*")
	expectedAlpha := "A"
    if alpha != expectedAlpha {
        t.Errorf("got %s but expected %s",alpha,expectedAlpha)
    }
}