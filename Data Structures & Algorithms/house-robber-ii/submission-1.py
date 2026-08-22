class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums)==1:
            return nums[0]
        if len(nums) == 2:
            return max(nums[0], nums[1])
        a = nums[0]
        b = max(a, nums[1])
        c = nums[1]
        d = max(c, nums[2])
        for i in range(2, len(nums)-1):
            a, b = b, max(a+nums[i], b)
        for i in range(3, len(nums)):
            c, d = d, max(c+nums[i], d)
        




        return max(b,d)
        