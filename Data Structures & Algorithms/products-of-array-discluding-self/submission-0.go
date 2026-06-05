func productExceptSelf(nums []int) []int {

	// simple soln : calc overall product o(n)
	// count zeros
	// if count =1: skip zero in product. 
	// while appending results, if zero count is =1, every non zero index has 0  and zero index has product
	// if count>1: append all zeros. 
	//  if count = 0
	// sing trvseral thru slice , divide nums[i] and append to result.

	count:=0
	product := 1
	res := make([]int, len(nums))
	for _, num := range nums{
		if num==0 {count++}
	} 

	if count>1{
		return res
	}
	
	for _, num:= range nums{
		if num!=0 {product *= num}
	}
	
	if count == 1{
		for i:=0;i<len(nums);i++{
			if nums[i]==0 {
				res[i]  = product
				}else {
					res[i] = 0
					}
		}
	}else{
		for i:=0;i<len(nums);i++{
			res[i] = product/nums[i]
		}
	}
	return res

}
