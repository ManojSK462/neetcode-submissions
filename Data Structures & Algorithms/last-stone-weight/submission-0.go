type minheap []int

func (h minheap) Len()int {return len(h)}

func(h minheap) Less(i, j int)bool {return h[i]>h[j]}

func(h minheap) Swap(i, j int){h[i], h[j]= h[j], h[i]}

func(h *minheap) Push(x any){*h = append(*h, x.(int))}

func(h *minheap) Pop() any {
	old:=*h
	p := old[len(old)-1]
	*h = old[:len(old)-1]
	return p
}


func lastStoneWeight(stones []int) int {
	h := minheap{}
	heap.Init(&h)

	for _, s := range stones{
		heap.Push(&h, s)
	}

	for h.Len()>1{
		s1 := heap.Pop(&h).(int)
		s2 := heap.Pop(&h).(int)

		if s1!=s2 {heap.Push(&h, s1-s2)}
	}

	if h.Len()==0{return 0}
	
	laststone:= heap.Pop(&h).(int)

	return laststone

}
