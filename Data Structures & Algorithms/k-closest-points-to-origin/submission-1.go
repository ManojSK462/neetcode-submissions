type item struct {
	dist  int
	index int
}

type minheap []item

func (h minheap) Len()int {return len(h)}

func(h minheap) Less(i, j int)bool {return h[i].dist<h[j].dist}

func(h minheap) Swap(i, j int){h[i], h[j]= h[j], h[i]}

func(h *minheap) Push(x any){*h = append(*h, x.(item))}

func(h *minheap) Pop() any {
	old:=*h
	p := old[len(old)-1]
	*h = old[:len(old)-1]
	return p
}


func kClosest(points [][]int, k int) [][]int {

	distheap := minheap{}

	for i:=0;i<len(points);i++ {
		dist := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		heap.Push(&distheap, item{dist: dist, index: i})
	}

	result := make([][]int, 0, k)

	for i := 0; i < k; i++ {
		current := heap.Pop(&distheap).(item)
		result = append(result, points[current.index])
	}

	return result

}
