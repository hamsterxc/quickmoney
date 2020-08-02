package main

type Market struct {
	stocks [][]Stock
}

func (this *Market) isAvailableToBuy() bool {
	return len(this.stocks) > 2
}

func (this *Market) getAvailable() (available []Stock) {
	if this.isAvailableToBuy() {
		for _, column := range this.stocks {
			available = append(available, column[len(column)-1])
		}
	}
	return
}

func (this *Market) take(index int) (result Stock) {
	var stocks [][]Stock

	for i, column := range this.stocks {
		if i == index {
			last := len(column) - 1
			result = column[last]
			column = column[:last]
		}

		if len(column) > 0 {
			stocks = append(stocks, column)
		}
	}
	this.stocks = stocks

	return
}
