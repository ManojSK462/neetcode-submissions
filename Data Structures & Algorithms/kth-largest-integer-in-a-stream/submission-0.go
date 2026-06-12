type minheap []int

func (h minheap) Len()int {return len(h)}

func(h minheap) Less(i, j int)bool {return h[i]<h[j]}

func(h minheap) Swap(i, j int){h[i], h[j]= h[j], h[i]}

func(h *minheap) Push(x any){*h = append(*h, x.(int))}

func(h *minheap) Pop() any {
	old:=*h
	p := old[len(old)-1]
	*h = old[:len(old)-1]
	return p
}


type KthLargest struct {
    h minheap
	k int
}


func Constructor(k int, nums []int) KthLargest {
    hp := KthLargest{k:k}

	for _, num := range nums{
		hp.Add(num)
	}

	return hp
}


func (this *KthLargest) Add(val int) int {
    heap.Push(&this.h, val)

	if this.h.Len()>this.k {
		heap.Pop(&this.h)
	}

	return this.h[0]

}
