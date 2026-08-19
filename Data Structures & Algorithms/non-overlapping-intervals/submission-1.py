class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda x:x[1])
        count, prev_end = 0, float('-inf')
        for i in intervals:
            if i[0]>=prev_end:
                count+=1
                prev_end = i[1]
        return len(intervals)-count
            
            
