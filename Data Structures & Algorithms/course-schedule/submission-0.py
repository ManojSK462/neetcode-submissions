class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        courses = defaultdict(list)
        pending = [0]*numCourses

        for p, c in prerequisites:
            courses[p].append(c)
            pending[c]+=1
        
        q = deque([i for i in range(numCourses) if pending[i]==0 ])

        while q:
            curr = q.popleft()
            for course in courses[curr]:
                pending[course]-=1
                if pending[course]==0:
                    q.append(course)

        for i in range(numCourses):
            if pending[i]!=0:
                return False
        return True
                    




        