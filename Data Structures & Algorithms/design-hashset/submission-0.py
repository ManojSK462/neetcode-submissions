class MyHashSet:

    def __init__(self):
        self.size = 1000
        self.data = [[] for i in range(self.size)]
        
    def add(self, key: int) -> None:
        b = self.data[key%1000]
        if key not in b:
            b.append(key)

    def remove(self, key: int) -> None:
        b = self.data[key%1000]
        if key in b:
            b.remove(key)
        
    def contains(self, key: int) -> bool:
        return key in self.data[key%1000]

        


# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)