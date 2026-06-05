func productExceptSelf(nums []int) []int {

	// simple soln : calc overall product o(n)
	// count zeros
	// if count =1: skip zero in product. 
	// while appending results, if zero count is =1, every non zero index has 0  and zero index has product
	// if count>1: append all zeros. 
	//  if count = 0
	// sing trvseral thru slice , divide nums[i] and append to result.

// 	

left := make([]int, len(nums))
right := make([]int , len(nums))
lp := 1
rp := 1

left[0] = 1

right[len(nums)-1] = 1

//left product pass
for i:=0;i<len(nums)-1;i++{
	lp *= nums[i]
	left[i+1] = lp
}
//right prodcut pass
for i:=len(nums)-1;i>0;i--{
	rp *= nums[i]
	right[i-1] = rp
}

//final result
res := make([]int , len(nums))

for i:=0;i<len(nums);i++{
	res[i] = left[i]*right[i]
}

return res
}
