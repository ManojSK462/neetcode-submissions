class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:

        intervals.sort(key = lambda x:x[0])
        res = [intervals[0]]
        for i in intervals[1:]:
            if i[0] <= res[-1][1]:
                end = max(i[1], res[-1][1])
                res[-1][1] = end
            else:
                res.append(i)
        return res

        