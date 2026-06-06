func trap(height []int) int {

lp := 0
rp := 0
n := len(height)
left_max_h := make([]int, n)
right_max_h := make([]int, n)

area:=0

for i:=0;i<n;i++{
	
	left_max_h[i] = lp
	lp = max(height[i], lp)

	right_max_h[n-i-1] = rp
	rp = max(height[n-i-1], rp)
}




for i:=0;i<n;i++ {
	water := min(left_max_h[i], right_max_h[i]) - height[i]
	if water>0{area += water}
	
}

return area
}
