func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	target:=0
	for i:=0;i<len(nums);i++{
		if i!=0 && nums[i]==nums[i-1]{continue}
		j := i+1
		k := len(nums)-1
	for j<k{
		if nums[i]+nums[j]+nums[k] == target {
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
		if nums[i]+nums[j]+nums[k] < target{
			j++
		}
		if nums[i]+nums[j]+nums[k] > target{
			k--
		}
		
	}
		
	}

return res
}
