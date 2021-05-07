package shared

import "math"

func RoundDecimal(num float64) float64 {
	return math.Round(num)
}

func RoundUp(num, places float64) float64 {
	shift := math.Pow(10, places)
	return RoundDecimal(num*shift) / shift
}

func CalcAmount(price, budget, minAmount, places float64) float64 {
	amount := RoundUp(budget/places, places)
	if amount > minAmount {
		return minAmount
	}
	return amount
}
