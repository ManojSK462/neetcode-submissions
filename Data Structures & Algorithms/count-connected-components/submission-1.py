class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        graph = defaultdict(list)

        for u,v in edges:
            graph[u].append(v)
            graph[v].append(u)

        
        visited = [False]*n
        components = 0
        q = deque()
        for i in range(n):
            if not visited[i]:
                components+=1
                visited[i] = True
                q.append(i)
            while q:
                node = q.popleft()
                for nei in graph[node]:
                    if not visited[nei]:
                        visited[nei] = True
                        q.append(nei)

        return components


        