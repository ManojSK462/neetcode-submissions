type DynamicArray struct {
a []int
size int
capacity int
}

func NewDynamicArray(capacity int) DynamicArray {
    return DynamicArray{
        a: make([]int, capacity),
        size: 0,
        capacity: capacity,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.a[i]

}

func (da *DynamicArray) Set(i int, n int) {
    da.a[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.size == da.capacity{
        da.resize()
    }
    da.a[da.size] = n
    da.size++
}

func (da *DynamicArray) Popback() int {
    da.size--
    return da.a[da.size]
}

func (da *DynamicArray) resize() {
    da.capacity*=2
    na := make([]int, da.capacity)
    for i:=0;i<da.size;i++{
        na[i] = da.a[i]
    }
    da.a = na
}

func (da *DynamicArray) GetSize() int {
    return da.size
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
