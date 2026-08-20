class Solution:
    def ispalin(self, s):
        i, j = 0, len(s)-1
        while i<j:
            if s[i]!=s[j]:
                return False
            i+=1
            j-=1
        return True
    def partition(self, s: str) -> List[List[str]]:
        res, part = [], []
        def dfs(i, j):
            if j >= len(s):
                if i==j:
                    res.append(part.copy())
                return
            if self.ispalin(s[i:j+1]):
                part.append(s[i:j+1])
                dfs(j+1, j+1)
                part.pop()
            dfs(i, j+1)
                
            
        dfs(0,0)
        return res
