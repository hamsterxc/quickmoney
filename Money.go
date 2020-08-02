package main

type Money struct {
	value int
}

func plus(first *Money, second *Money) Money {
	return Money{value: first.value + second.value}
}

func minus(first *Money, second *Money) (Money, bool) {
	if first.value >= second.value {
		return Money{value: first.value - second.value}, true
	} else {
		return Money{}, false
	}
}

func mult(first *Money, second *Money) Money {
	return Money{value: first.value * second.value}
}
