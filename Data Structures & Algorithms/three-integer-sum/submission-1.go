func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	target:=0
	for i:=0;i<len(nums);i++{
		if i!=0 && nums[i]==nums[i-1]{continue}
		j := i+1
		k := len(nums)-1
	for j<k{
		sum:=nums[i]+nums[j]+nums[k]
		if sum == target {
			res = append(res, []int{nums[i],nums[j],nums[k]})
			j++
			k--
			for j<k && nums[j]==nums[j-1]{
				j++
			}
			for j<k && nums[k]==nums[k+1]{
				k--
			}
		}
		if sum < target{
			j++
		}
		if sum > target{
			k--
		}
		
	}
		
	}

return res
}
