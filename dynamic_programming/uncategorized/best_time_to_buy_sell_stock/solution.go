package solution

import "math"

func BottomUp(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	cheapestPricesOnDay := make([]int, len(prices))
	for i, price := range prices {
		if i == 0 {
			cheapestPricesOnDay[i] = price
		} else {
			cheapestPricesOnDay[i] = int(math.Min(
				float64(cheapestPricesOnDay[i-1]),
				float64(price),
			))
		}
	}

	var maxProfit float64 = 0
	for i := 1; i < len(prices); i++ {
		buyPrice := cheapestPricesOnDay[i-1]
		sellPrice := prices[i]

		profit := sellPrice - buyPrice

		maxProfit = math.Max(maxProfit, float64(profit))
	}

	return int(maxProfit)
}
