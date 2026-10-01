class Solution:
    def minInterval(self, intervals: List[List[int]], queries: List[int]) -> List[int]:

        intervals = sorted(intervals, key=lambda x: x[0])
        order = sorted(range(len(queries)), key=lambda k: queries[k])

        res = [-1]*len(queries)
        i=0
        n = len(intervals)
        heap = []
        for k in order:
            q = queries[k]
            while i<n and intervals[i][0]<=q:
                heapq.heappush(heap, (intervals[i][1]-intervals[i][0]+1, intervals[i][1]))
                i+=1
            while heap and heap[0][1]<q:
                heapq.heappop(heap)
            if heap:
                res[k] = heap[0][0]

        return res


        



       