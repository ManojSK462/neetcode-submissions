func longestConsecutive(nums []int) int {

seen := make(map[int]bool)
for i:=0;i<len(nums);i++{
	seen[nums[i]] = true
}

best:=0
for num:= range seen{
	if !seen[num-1]{
		length:=1
		curr := num
		for seen[curr+1]{
			curr++
			length++
		}

		if length>best{best = length}
	}
}
return best

}
