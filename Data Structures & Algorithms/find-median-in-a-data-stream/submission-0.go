type intheap struct{
	nums []int
	isMin bool
}

func (h intheap) Len()int {return len(h.nums)}

func(h intheap) Less(i, j int)bool {
	if h.isMin {
		return h.nums[i]<h.nums[j]
		}else{
			return h.nums[i]>h.nums[j]
		}
}

func(h intheap) Swap(i, j int){h.nums[i], h.nums[j]= h.nums[j], h.nums[i]}

func(h *intheap) Push(x any){h.nums = append(h.nums, x.(int))}

func(h *intheap) Pop() any {
	old:=h.nums
	p := old[len(old)-1]
	h.nums = old[:len(old)-1]
	return p
}


type MedianFinder struct {
    minheap intheap
	maxheap intheap
}


func Constructor() MedianFinder {
    return MedianFinder{
		minheap: intheap{isMin: true},
		maxheap: intheap{isMin: false},
	}
}


func (this *MedianFinder) AddNum(num int)  {
    if this.maxheap.Len()==0 || num <= this.maxheap.nums[0]{
		heap.Push(&this.maxheap, num)
	}else{
		heap.Push(&this.minheap, num)
	}

	//balance now. maxheap is allowed to have +1 element. that is where we fecth median from
	if this.maxheap.Len() - this.minheap.Len() > 1{
		x:= heap.Pop(&this.maxheap)
		heap.Push(&this.minheap, x)
	}
	if this.minheap.Len() - this.maxheap.Len() > 0{
		x:= heap.Pop(&this.minheap)
		heap.Push(&this.maxheap, x)
	}

}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxheap.Len() == this.minheap.Len(){
		return (float64(this.maxheap.nums[0]) + float64(this.minheap.nums[0]))/2.0
	}else{
		return float64(this.maxheap.nums[0])  
	}
}
