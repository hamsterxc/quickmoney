package main

import (
	"fmt"
	"strings"
)

func (this *Game) String() (result string) {
	result = this.gameState.market.String()

	for i, playerName := range this.playerNames {
		result += "\n"
		player := this.gameState.players[i]

		if i == this.turn {
			result += "-> "
		}
		result += fmt.Sprintf("%s - $%d", playerName, player.money.value)
		result += "\n"

		if len(player.stocks) > 0 {
			for _, stock := range player.stocks {
				result += fmt.Sprintf("%s  ", stock.String())
			}
			result += "\n"

			if i == this.turn {
				for j := range player.stocks {
					space := ""
					if j < 9 {
						space = " "
					}
					result += fmt.Sprintf("(%d)%s", j+1, space)
				}
				result += "\n"
			}
		} else {
			result += "(No stocks owned)\n"
		}
	}

	result += fmt.Sprintf("\n%s's turn!\n", this.playerNames[this.turn])

	return
}

func (this *Market) String() (result string) {
	result = "Stock Market\n"

	maxColumns := len(this.stocks)
	maxRows := 0
	for _, column := range this.stocks {
		if len(column) > maxRows {
			maxRows = len(column)
		}
	}

	// transposed matrix of stock string representations, either "  " or "VC"
	stockStrings := make([][]string, maxRows)
	for i := 0; i < maxRows; i++ {
		stockStrings[i] = make([]string, maxColumns)
		for j := 0; j < maxColumns; j++ {
			if i >= len(this.stocks[j]) {
				stockStrings[i][j] = "  "
			} else {
				stock := this.stocks[j][i]
				stockStrings[i][j] = stock.String()
			}
		}
	}

	// top border
	result += "+"
	for i := 0; i < maxColumns; i++ {
		result += "----+"
	}
	result += "\n"

	// stocks
	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxColumns; j++ {
			s := stockStrings[i][j]
			result += choose(!isEmpty(s) && ((j == 0) || isEmpty(stockStrings[i][j-1])), "|", choose(j == 0, " ", ""))
			result += fmt.Sprintf(" %s ", s)
			result += choose(!isEmpty(s), "|", choose((j < maxColumns-1) && isEmpty(stockStrings[i][j+1]), " ", ""))
		}
		result += "\n"
		for j := 0; j < maxColumns; j++ {
			s := stockStrings[i][j]
			result += choose(!isEmpty(s) && ((j == 0) || isEmpty(stockStrings[i][j-1])), "+", choose(j == 0, " ", ""))
			result += choose(isEmpty(s), "    ", "----")
			result += choose(!isEmpty(s), "+", choose((j < maxColumns-1) && isEmpty(stockStrings[i][j+1]), " ", ""))
		}
		result += "\n"
	}

	// index helpers
	for i := 0; i < maxColumns; i++ {
		result += fmt.Sprintf(" (%d) ", i+1)
	}
	result += "\n"

	return
}

var stockColorCode = map[StockColor]string{
	JOKER:  "J",
	RED:    "R",
	GREEN:  "G",
	BLUE:   "B",
	YELLOW: "Y",
	PURPLE: "P",
}

func (this *Stock) String() string {
	return fmt.Sprintf("%d%s", this.value.value, stockColorCode[this.color])
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func choose(condition bool, ifTrue string, ifFalse string) string {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}
