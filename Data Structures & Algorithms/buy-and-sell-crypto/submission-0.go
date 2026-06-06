func maxProfit(prices []int) int {

	best_cp := prices[0]
	best_profit := 0

	for i:=1;i<len(prices);i++ {
		best_profit = max(best_profit, prices[i] - best_cp)
		best_cp = min(best_cp, prices[i])
	} 

	return best_profit
}
