class Solution:
    def combinationSum2(self, nums: List[int], target: int) -> List[List[int]]:
        res = []
        nums.sort()
        def dfs(i, curr, currsum):
            if currsum>=target or i>=len(nums):
                if currsum == target:
                    res.append(curr.copy())
                return
            if currsum+nums[i]<=target:
                curr.append(nums[i])
                currsum+=nums[i]
                dfs(i+1, curr, currsum)
                curr.pop()
                currsum -= nums[i]
            
            x = 1
            for j in range(i+1, len(nums)):
                if nums[j]!=nums[i]:
                    break
                x+=1
            
            dfs(i+x, curr, currsum)

        dfs(0, [], 0)
        return res

            
            

