class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        i = 0
        window = set()
        for j in range(len(nums)):
            if abs(i-j) > k:
                window.remove(nums[i])
                i+=1
            if nums[j] in window:
                return True
            window.add(nums[j])
            
            
        return False

        #i:0 j:1 k = 1 window: 2,1

            
            
            
                

        