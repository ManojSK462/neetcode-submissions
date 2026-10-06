class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        max_freq = 0
        window = {}
        i = 0
        res = 0

        for j in range(len(s)):
            window[s[j]] = window.get(s[j], 0) + 1
            max_freq = max(max_freq, window[s[j]])
            if j-i+1 - max_freq > k:
                window[s[i]]-=1
                i+=1
            res = max(res, j-i+1)

        return res
        