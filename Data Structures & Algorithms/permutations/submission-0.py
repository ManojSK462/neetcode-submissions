class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        track = set()
        def dfs(curr):
            if len(curr) == len(nums):
                res.append(curr.copy())
                return
            for i in nums:
                if i not in track:
                    curr.append(i)
                    track.add(i)
                    dfs(curr)
                    curr.pop()
                    track.remove(i)

        dfs([])
        return res
        