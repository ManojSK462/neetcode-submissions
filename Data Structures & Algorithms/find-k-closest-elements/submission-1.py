class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        window = []
        for j in range(len(arr)):
            if len(window)==k:
                if abs(arr[j]-x)<abs(window[0]-x):
                    window = window[1:]
                elif abs(arr[j]-x)==abs(window[0]-x):
                    continue
                else:
                    break
            window.append(arr[j]) #window: 1
        return window


        