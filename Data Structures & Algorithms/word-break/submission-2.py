class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        dp = [False]*(len(s)+1)
        dp[0] = True 

        for i in range(len(s)+1):
            for word in wordDict:
                if (i-len(word))>=0 and s[i-len(word):i] == word:
                    dp[i] = dp[i-len(word)]
                if dp[i]:
                    break
        return dp[len(s)]


        