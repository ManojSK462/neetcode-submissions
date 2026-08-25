class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:

        def reachable(heights, startQ, r, c):
            visited = set()
            for pair in startQ:
                visited.add(pair)
            while startQ:
                x, y = startQ.popleft()
                for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                    nx, ny = x+dx , y+dy
                    if 0<=nx<r and 0<=ny<c and (nx, ny) not in visited and heights[x][y]<=heights[nx][ny]:
                        visited.add((nx, ny))
                        startQ.append((nx, ny))
            return visited
        

        r, c = len(heights), len(heights[0])

        pacificQ, atlanticQ = deque(), deque()

        for i in range(r):
            pacificQ.append((i, 0))
            atlanticQ.append((i, c-1))
        for i in range(c):
            pacificQ.append((0, i))
            atlanticQ.append((r-1, i))

        result = reachable(heights, pacificQ, r, c) & reachable(heights, atlanticQ, r, c)
        result = [[r,c] for r,c in result]
        return result
                
        
                       


        