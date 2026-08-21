class Solution:
	def maxProfit(self, prices: List[int]) -> int:
		cost_price = prices[0]
		profit = 0
		for stock in prices[1:]:
			profit = max(profit, stock - cost_price)
			cost_price = min(cost_price, stock)

		return profit
        