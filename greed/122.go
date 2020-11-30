package greed

func maxProfit(prices []int) int {
	allMaxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			allMaxProfit += prices[i] - prices[i-1]
		}
	}
	return allMaxProfit
}
