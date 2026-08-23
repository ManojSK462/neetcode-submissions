class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [float('inf') for i in range(amount+1)]
        dp[0] = 0
        for i in range(1, amount+1):
            ways = float('inf')
            for coin in coins:
                if amount>=coin:
                    ways = min(ways, dp[i-coin])
            dp[i] = 1 + ways
        return -1 if dp[amount]==float('inf') else dp[amount]



        