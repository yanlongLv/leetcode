package greed

func maxProfit(prices []int, fee int) int {
	result := 0
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i] >= minPrice && prices[i] <= minPrice+fee {
			continue
		}
		if prices[i] > minPrice+fee {
			result += prices[i] - minPrice - fee
			minPrice = prices[i] - fee
		}
	}
	return result
}
