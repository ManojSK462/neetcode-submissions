func twoSum(numbers []int, target int) []int {
	left :=0
	right:=len(numbers)-1

	for left<right {
		sum := numbers[left] + numbers[right]
		if sum==target {return []int{left+1, right+1}}
		if sum<target {left++}
		if sum>target  {right--}
		
	}
	return []int{-1, -1}
}
