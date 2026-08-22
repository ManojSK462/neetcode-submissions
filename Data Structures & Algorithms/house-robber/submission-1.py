class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums)==1:
            return nums[0]
        a = nums[0]
        b = max(a, nums[1])
        for i in range(2, len(nums)):
            a, b = b, max(a+nums[i], b)
        return b
        