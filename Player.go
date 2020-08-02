package main

type Player struct {
	stocks []Stock
	money  Money
}

func (this *Player) buy(stock *Stock) bool {
	buyValue := stock.getBuyValue()
	newMoney, success := minus(&this.money, &buyValue)
	if success {
		this.money = newMoney
		this.stocks = append(this.stocks, *stock)
	}
	return success
}

func (this *Player) sell(first int, second int) bool {
	stockFirst := this.stocks[first]
	stockSecond := this.stocks[second]
	if !stockFirst.isCompatible(&stockSecond) {
		return false
	}

	if first < second {
		first, second = second, first
	}
	this.stocks = remove(this.stocks, first)
	this.stocks = remove(this.stocks, second)

	sellValue := stockFirst.getSellValue(&stockSecond)
	this.money = plus(&this.money, &sellValue)

	return true
}

func remove(stocks []Stock, index int) []Stock {
	count := len(stocks)
	stocks[index] = stocks[count-1]
	return stocks[:count-1]
}
