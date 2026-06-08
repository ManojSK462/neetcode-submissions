type MinStack struct {

stack []int
minstack []int

}

func Constructor() *MinStack {

	return &MinStack{
		stack: []int{},
		minstack: []int{},
		
	}

}

// func (this *MinStack) Push(val int) {
// 	this.stack = append(this.stack, val)

// 	if len(this.minstack)>0{
// 		if val<this.minstack[len(this.minstack)-1]{
// 			this.minstack = append(this.minstack, val)
// 			}		
// 	}else{
// 		this.minstack = append(this.minstack, val)
// 	}
// }


// func (this *MinStack) Pop() {
// 	poppedVal := this.stack[len(this.stack)-1]
// 	this.stack = this.stack[:len(this.stack)-1]

// 	if poppedVal == this.minstack[len(this.minstack)-1] {
// 		this.minstack = this.minstack[:len(this.minstack)-1]
// 	}
// }

// func (this *MinStack) Top() int {
	
// 		return this.stack[len(this.stack)-1]
	
// }

// func (this *MinStack) GetMin() int {
	
// 		return this.minstack[len(this.minstack)-1]
	
// }

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minstack) == 0 || val <= this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	}
}

func (this *MinStack) Pop() {
	poppedVal := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if poppedVal == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
