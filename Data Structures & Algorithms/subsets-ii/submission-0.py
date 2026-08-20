class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()

        res = []
        def dfs(i, curr):
            if i>=len(nums):
                res.append(curr.copy())
                return
            
            curr.append(nums[i])
            dfs(i+1, curr)
            curr.pop()
            x = 1
            for j in range(i+1, len(nums)):
                
                if nums[j]!=nums[i]:
                    break
                x+=1
            if x!=0:
                dfs(i+x, curr)
        
        dfs(0, [])
        return res
                


        