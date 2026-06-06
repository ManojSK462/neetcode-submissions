func maxArea(heights []int) int {
	max:=0
	area:=0

	l:=0
	r:=len(heights)-1

	for l<r {
		area = min(heights[l], heights[r])*(r-l)
		if area>max {max = area}
		if heights[l]<heights[r]{
			l++
		}else{
			r--
		}
	}
	return max
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}