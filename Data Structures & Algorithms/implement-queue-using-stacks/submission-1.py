class MyQueue:

    def __init__(self):
        self.inb = []
        self.out = []
        

    def push(self, x: int) -> None:
        self.inb.append(x)
        
    def move(self):
        if not self.out:
            while self.inb: self.out.append(self.inb.pop())

    def pop(self) -> int:
        self.move()
        return self.out.pop()
        

    def peek(self) -> int:
        self.move()
        return self.out[-1]
        

    def empty(self) -> bool:
        return len(self.inb)==0 and len(self.out)==0
        


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()