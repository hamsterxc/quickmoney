package main

type Stock struct {
	color StockColor
	value Money
}

type StockColor int

const (
	JOKER StockColor = iota
	RED
	GREEN
	BLUE
	YELLOW
	PURPLE
)

func (this *Stock) isCompatible(stock *Stock) bool {
	return (this.color == JOKER) || (stock.color == JOKER) || (this.color == stock.color)
}

func (this *Stock) getBuyValue() Money {
	return this.value
}

func (this *Stock) getSellValue(stock *Stock) Money {
	return mult(&this.value, &stock.value)
}
