from collections import deque
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int: 
        r, c = len(grid), len(grid[0])
        seen = [[False]*c for i in range(r)]
        count = 0
        for i in range(r):
            for j in range(c):
                if grid[i][j] == '1' and not seen[i][j]:
                    count+=1
                    q = deque([(i,j)])
                    seen[i][j] = True
                    while q:
                        x, y = q.popleft()
                        for dx, dy in ((1,0), (-1,0), (0,1), (0,-1)):
                            nx, ny = x+dx, y+dy
                            if 0<=nx<r and 0<=ny<c and grid[nx][ny] == '1' and not seen[nx][ny]:
                                seen[nx][ny] = True
                                q.append((nx, ny))
        return count



