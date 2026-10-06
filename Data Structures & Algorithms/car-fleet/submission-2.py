class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        pairs = sorted(zip(position, speed), reverse=True)

        fleets = 0

        worsetime = 0.0


        for _, i in pairs:
            time = (target - _)/(i)
            if time>worsetime:
                fleets+=1
                worsetime = time
        
        return fleets
        