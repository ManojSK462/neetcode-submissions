class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        d = {}
        stack = []

        for n in nums2:
            while stack and stack[-1]<n:
                d[stack.pop()] = n
            stack.append(n)

        return [d.get(i, -1) for i in nums1]

        