func findDuplicate(nums []int) int {
    
	for i:=0;i<len(nums);i++{
		idx := abs(nums[i])-1

		if nums[idx]<0{
			return abs(nums[i])
		}
		nums[idx]*=-1
	}

return -1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
