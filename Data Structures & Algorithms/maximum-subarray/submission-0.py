class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_ = sum_ = nums[0]
        for n in nums[1:]:
            sum_ = max(n, sum_+n)
            max_ = max(max_, sum_)
        return max_
        