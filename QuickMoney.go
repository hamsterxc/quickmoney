package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var DELIMITER = strings.Repeat("=", 40)

type TurnState int

const (
	START TurnState = iota
	BUYING
	SELLING
	FINISH
)

func main() {
	rand.Seed(time.Now().UnixNano())

	playersCount := readPlayersCount()
	playerNames := readPlayerNames(playersCount)
	game := newGame(playerNames)

	for !game.isFinished() {
		fmt.Printf("\n%s\n\n%s\n", DELIMITER, game.String())

		isMarketClosed := !game.gameState.market.isAvailableToBuy()
		if isMarketClosed {
			fmt.Print("End game phase - stock market is closed for buyers!\n\n")
		}

		playerName := game.playerNames[game.turn]
		player := &game.gameState.players[game.turn]
		state := START
		for state != FINISH {
			switch state {
			case START:
				prompt := fmt.Sprintf("%s, your action ((1) - Buy, (2) - Sell): ", playerName)
				action := readOneInt(prompt, func(value int) bool {
					return (value == 1) || (value == 2)
				})

				switch action {
				case 1:
					state = BUYING
				case 2:
					state = SELLING
				}

			case BUYING:
				if isMarketClosed {
					fmt.Println("Stock market is closed for buyers!")
					state = START
					continue
				}

				stocks := game.gameState.market.getAvailable()
				count := len(stocks)
				prompt := fmt.Sprintf("%s, choose a stock to buy (1..%d, enter none to return): ", playerName, count)
				index := readOneIntDefault(prompt, func(value int) bool {
					return (value >= 1) && (value <= count)
				}, 0)

				if index == 0 {
					state = START
				} else {
					index--
					success := player.buy(&stocks[index])
					if success {
						game.gameState.market.take(index)
						state = FINISH
					} else {
						fmt.Println("You don't have enough money!")
						state = START
					}
				}

			case SELLING:
				stocks := player.stocks
				count := len(stocks)
				if count < 2 {
					fmt.Println("Not enough stocks to sell!")
					state = START
					continue
				}

				prompt := fmt.Sprintf("%s, choose two different stocks to sell (1..%d, enter none to return): ", playerName, count)
				first, second := readTwoIntsDefault(prompt, func(first int, second int) bool {
					return (first >= 1) && (first <= count) && (second >= 1) && (second <= count) && (first != second)
				}, 0, 0)

				if (first == 0) && (second == 0) {
					state = START
				} else {
					first--
					second--
					if stocks[first].isCompatible(&stocks[second]) {
						player.sell(first, second)
						state = FINISH
					} else {
						fmt.Println("Stocks should be of the same color, or at least one of the should be a Joker!")
						state = START
					}
				}

				if isMarketClosed {
					game.playerFinalSell[game.turn] = true
				}
			}
		}

		game.turn = (game.turn + 1) % len(game.gameState.players)
	}

	fmt.Printf("\n%s\n\nFinal standings\n\n", DELIMITER)
	winnerNames := make([]string, 0)
	winnerMoney := 0
	for index, playerName := range game.playerNames {
		money := game.gameState.players[index].money.value
		fmt.Printf("%s has $%d\n", playerName, money)
		switch {
		case money > winnerMoney:
			winnerNames = []string{playerName}
			winnerMoney = money
		case money == winnerMoney:
			winnerNames = append(winnerNames, playerName)
		}
	}

	fmt.Println()
	switch {
	case len(winnerNames) == 1:
		fmt.Printf("%s won!\n", winnerNames[0])
	case len(winnerNames) > 1:
		fmt.Printf("%s tied!\n", strings.Join(winnerNames, ", "))
	}
}

func readPlayersCount() int {
	return readOneIntDefault("Number of players playing (2..4) [2]: ", func(value int) bool {
		return (value >= 2) && (value <= 4)
	}, 2)
}

func readPlayerNames(playersCount int) []string {
	ordinals := map[int]string{
		0: "First",
		1: "Second",
		2: "Third",
		3: "Fourth",
	}

	playerNames := make([]string, playersCount)
	for i := 0; i < playersCount; i++ {
		var ordinal string
		if i < len(ordinals) {
			ordinal = ordinals[i]
		} else {
			ordinal = fmt.Sprintf("%dth", i+1)
		}

		var def = fmt.Sprintf("Player %d", i+1)
		playerNames[i] = readOneStringDefault(
			fmt.Sprintf("%s player's name [%s]: ", ordinal, def),
			func(value string) bool { return true },
			def)
	}

	return playerNames
}
