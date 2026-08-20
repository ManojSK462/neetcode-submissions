class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        pattern = []
        def dfs(i, j):
            if len(pattern) == 2*n:
                res.append("".join(pattern))
                return
            if i<n:
                pattern.append("(")
                dfs(i+1, j)
                pattern.pop()
            if j<i:
                pattern.append(")")
                dfs(i, j+1)
                pattern.pop()
            
        
        dfs(0 , 0)
        return res
            

        
       