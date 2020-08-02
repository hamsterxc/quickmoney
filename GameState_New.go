package main

import (
	"math/rand"
)

func newGameState(playersCount int) GameState {
	players := make([]Player, playersCount)
	for i := 0; i < playersCount; i++ {
		players[i] = newPlayer()
	}

	return GameState{
		players: players,
		market:  newMarket(),
	}
}

func newPlayer() Player {
	joker := Stock{
		color: JOKER,
		value: Money{value: 2},
	}
	return Player{
		stocks: []Stock{joker},
		money:  Money{value: 20},
	}
}

func newMarket() Market {
	values := []int{2, 2, 3, 3, 4, 5, 6}
	colors := []StockColor{RED, GREEN, BLUE, YELLOW, PURPLE}
	columns := 5
	rows := 7

	stocks := make([]Stock, len(values)*len(colors))
	index := 0
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(colors); j++ {
			stocks[index] = Stock{
				color: colors[j],
				value: Money{value: values[i]},
			}
			index++
		}
	}
	rand.Shuffle(len(stocks), func(i, j int) {
		stocks[i], stocks[j] = stocks[j], stocks[i]
	})

	marketStocks := make([][]Stock, columns)
	index = 0
	for i := 0; i < columns; i++ {
		marketStocks[i] = make([]Stock, rows)
		for j := 0; j < rows; j++ {
			marketStocks[i][j] = stocks[index]
			index++
		}
	}

	return Market{stocks: marketStocks}
}
