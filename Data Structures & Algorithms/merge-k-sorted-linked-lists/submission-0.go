/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type minheap []*ListNode

func (h minheap) Len() int{
	return len(h)
}

func (h minheap) Less(i,j int) bool{
	return h[i].Val < h[j].Val
}

func(h minheap) Swap(i,j int){
	h[i], h[j] = h[j], h[i]
}

func(h *minheap) Push(x any){
	*h = append(*h, x.(*ListNode))
}
func(h *minheap) Pop() any{
	old := *h
	n := h.Len()

	node:=old[n-1]

	*h = old[:n-1]

	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
    
	h:=minheap{}

	heap.Init(&h)

	dummy:= &ListNode{}

	curr := dummy


	for i:=0;i<len(lists);i++{
		if lists[i]!=nil{
			heap.Push(&h, lists[i])
		} //push all heads
		
	}

	for h.Len()>0 {

		node := heap.Pop(&h).(*ListNode)

		if node.Next!=nil{
			heap.Push(&h, node.Next)
		}
		

		curr.Next = node

		curr = curr.Next

	}


	return dummy.Next

}
