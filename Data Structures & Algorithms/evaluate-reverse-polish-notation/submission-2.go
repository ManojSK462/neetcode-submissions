func evalRPN(tokens []string) int {

stack:= []int{}

for _, token := range tokens{

	if token=="+" || token=="-" || token=="*" || token=="/"{
		op1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		op2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		var res int

		if token=="+" {res= op1 + op2}
		if token=="-" {res= op2 - op1}
		if token=="*" {res= op1 * op2}
		if token=="/" {res= op2 / op1}

		stack = append(stack, res)

	}else{
		num, _:=strconv.Atoi(token)
		stack = append(stack, num )
	}


}


	return stack[0]

}
