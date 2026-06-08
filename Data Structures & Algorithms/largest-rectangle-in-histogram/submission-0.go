func largestRectangleArea(h []int) int {

stack := []int{}

left := make([]int, len(h))
right := make([]int, len(h))

for i:=0;i<len(h);i++{
	left[i] = -1
}


for i:=0;i<len(h);i++{
	right[i] = len(h)
}



for i:=0;i<len(h);i++{
	for len(stack)>0 && h[i]<h[stack[len(stack)-1]]{
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		right[idx] = i
	}
	stack  = append(stack, i)
}
for i:=len(h)-1;i>=0;i--{
	for len(stack)>0 && h[i]<h[stack[len(stack)-1]]{
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left[idx] = i
	}
	stack  = append(stack, i)
}


max_:=0

for i:=0;i<len(h);i++{

	area := h[i]*(right[i]-left[i]-1)
	
	if area>max_{max_=area}

}

return max_

}
