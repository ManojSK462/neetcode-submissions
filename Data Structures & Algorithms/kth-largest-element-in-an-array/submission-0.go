type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := minHeap{}

	for _, num := range nums {
		heap.Push(&h, num)

		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	return h[0]
}