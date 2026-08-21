class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        denoms = {5:0, 10:0, 20:0}
        for i in bills:
            denoms[i]+=1
            if i==10:
                if denoms[5]==0:
                    return False
                denoms[5]-=1
            elif i==20:
                if denoms[5]==0:
                    return False
                else:
                    if denoms[10]==0:
                        if denoms[5]>=3:
                            denoms[5]-=3
                        else:
                            return False
                    else:
                        denoms[5]-=1
                        denoms[10]-=1
            print(denoms)

        return True

        