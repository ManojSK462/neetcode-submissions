class Solution:
    def foreignDictionary(self, word: List[str]) -> str:

        graph = defaultdict(list)
        vertices = set(c for w in word for c in w)
        indegree = {v:0 for v in vertices}

        for i in range(1, len(word)):
            for j in range(min(len(word[i]), len(word[i-1]))):
                if word[i][j] != word[i-1][j]:
                    graph[word[i-1][j]].append(word[i][j])
                    indegree[word[i][j]]+=1
                    break
            else:
                if len(word[i-1])>len(word[i]):
                    return ""

        q = deque(i for i in vertices if indegree[i]==0)
        result = ""
        while q:
            curr = q.popleft()
            result+=curr
            for child in graph[curr]:
                indegree[child]-=1
                if indegree[child]==0:
                    q.append(child)
                    
        return result if len(result)==len(vertices) else ""
        


        