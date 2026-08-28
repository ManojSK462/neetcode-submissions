class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        d = {"(":")", "[":"]", "{":"}"}
        for c in s:
            if c=="(" or c=="[" or c=="{":
                stack.append(c)
            else:
                if stack:
                    top = stack.pop()
                    if not d[top] == c:
                        return False
                else:
                    return False
        return len(stack)==0
                
        

        