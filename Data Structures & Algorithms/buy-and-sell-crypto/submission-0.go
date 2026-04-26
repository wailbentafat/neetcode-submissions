func maxProfit(prices []int) int {
	max := 0
	min := math.MaxInt
	for i := 0 ; i<len(prices);i++{
		if min > prices[i]{
			min = prices[i]
		}
		if prices[i] - min > max{
			max = prices[i] - min
		}
		println(max)
	}
	return max
}

// func maxProfit(prices []int) int {
// 	minPrice := math.MaxInt
// 	maxProfit := 0
// 	for _ , v := range prices {
// 		if v < minPrice {
// 			minPrice = v
// 		}
// 		if v - minPrice > maxProfit {
// 			maxProfit = v - minPrice
// 		}
// 	}
// 	return maxProfit
// }