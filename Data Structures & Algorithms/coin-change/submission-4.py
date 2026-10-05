class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:

        dp = [float('inf')]*(amount+1)

        dp[0] = 0

        for target in range(1, amount+1):
            for denom in coins:
                if target-denom>=0:
                    dp[target] = min(dp[target], dp[target-denom]+1)
        return -1 if dp[-1]==float('inf') else dp[-1]
                

        