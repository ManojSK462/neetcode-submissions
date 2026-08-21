class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        counts = Counter(s)
        q = set()
        res = []
        substr_length = 0
        for c in s:
            substr_length+=1
            if counts[c]>0:
                q.add(c)
                counts[c]-=1
            if counts[c] == 0:
                q.remove(c)
            if len(q) == 0:
                res.append(substr_length)
                substr_length = 0
        return res

        # a:0, b:0, c:0
        #q : , , 
        # res : 6
