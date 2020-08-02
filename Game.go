package main

import (
	"math/rand"
)

type Game struct {
	playerNames     []string
	gameState       GameState
	turn            int
	playerFinalSell []bool
}

func newGame(playerNames []string) Game {
	playersCount := len(playerNames)
	return Game{
		playerNames:     playerNames,
		gameState:       newGameState(playersCount),
		turn:            rand.Intn(playersCount),
		playerFinalSell: make([]bool, playersCount),
	}
}

func (this *Game) isFinished() (finished bool) {
	finished = true
	for _, playerFinalSell := range this.playerFinalSell {
		finished = finished && playerFinalSell
	}
	return
}
