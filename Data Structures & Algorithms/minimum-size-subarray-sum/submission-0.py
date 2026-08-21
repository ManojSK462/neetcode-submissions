class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        i = 0
        subsum = 0
        minlen = float('inf')
        for j in range(len(nums)):
            subsum+=nums[j] # j=4, i=3, subsum = 6, minlen : 2
            if subsum>=target:
                minlen = min(minlen, j-i+1)
                while subsum>=target:
                    minlen = min(minlen, j - i + 1)
                    subsum-=nums[i]
                    i+=1
                    
                    
        return minlen if minlen != float('inf') else 0
                    



