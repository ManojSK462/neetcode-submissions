class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        for c in tokens:
            if c in ['+', '-', '*', '/']:
                op1 = stack.pop()
                op2 = stack.pop()

                if c == '+': stack.append(op1+op2)
                elif c == '-': stack.append(op2-op1)
                elif c=='*': stack.append(op1*op2)
                else: stack.append(int(op2/op1))

            else:
                stack.append(int(c))
        return stack[0]

