class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        window = {}
        res = 0
        i = 0
        maxf = 0
        for j in range(len(s)):
            window[s[j]] = window.get(s[j], 0) + 1
            maxf = max(maxf, window[s[j]])

            while j-i+1 > k+maxf:
                window[s[i]]-=1
                i+=1
            res = max(res, j-i+1)

        return res