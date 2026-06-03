// import "container/heap"

type Edge struct{
	to int
	weight int
}
type Item struct{
	dist int
	node int
}

type MinHeap []Item

func (h MinHeap) Len() int{
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any){
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any{
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}


func shortestPath(n int, edges [][]int, src int) map[int]int {
graph := make([][] Edge, n)

for _, edge := range edges{
	u := edge[0]
	v := edge[1]
	w := edge[2]

		graph[u] = append(graph[u], Edge{to:v, weight: w})
	}

	const INF = int(1e9)
	dist := make([]int, n)
	for i:=0; i<n; i++ {
		dist[i] = INF
	}
	dist[src] = 0

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Item{dist: 0, node: src})

	for h.Len()>0 {
		curr := heap.Pop(h).(Item)
		d := curr.dist
		node := curr.node

		if d > dist[node]{continue}

		for _, edge := range graph[node] {
			next := edge.to
			newDist := d + edge.weight

			if newDist < dist[next] {
				dist[next] = newDist
				heap.Push(h ,Item{dist: newDist, node:next})
			}
		}
	}
	res := make(map[int]int)

	for i := 0; i < n; i++ {
		if dist[i] == INF {
			res[i] = -1
		} else {
			res[i] = dist[i]
		}
	}

	return res

}
