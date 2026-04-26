func maxProfit(prices []int) int {
	
	maxProfit := 0;
	min := prices[0];

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		profit := prices[i] - min

		if (profit > maxProfit) {
			maxProfit = profit
		}
	}

	return maxProfit
}