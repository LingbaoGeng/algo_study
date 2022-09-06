package algo

import "math"

func GetPrice(prices []int) int {
	maxPrice := math.MinInt
	maxGap := math.MinInt

	for i := 0; i < len(prices); i++ {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else if maxPrice-prices[i] > maxGap {
			maxGap = maxPrice - prices[i]
		}
	}
	return maxGap
}
