// func leastInterval(tasks []byte, n int) int {
// 	count := make([]int, 26)
//     for _, task := range tasks {
//         count[task-'A']++
//     }


// 	maxf := 0
//     for _, cnt := range count {
//         if cnt > maxf {
//             maxf = cnt
//         }
//     }

// 	maxCount := 0
//     for _, cnt := range count {
//         if cnt == maxf {
//             maxCount++
//         }
//     }

// 	time := (maxf - 1) * (n + 1) + maxCount
//     if len(tasks) > time {
//         return len(tasks)
//     }
//     return time

// }
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	count := make(map[byte]int)
    for _, task := range tasks {
        count[task]++
    }

	h := &maxHeap{}

	for _, frequency := range count {
		heap.Push(h, frequency)
	}

	cooldown := make([][2]int, 0)
	time := 0

	 for h.Len() > 0 || len(cooldown) > 0 {
		time++
		if h.Len() > 0 {
			remaining := heap.Pop(h).(int) - 1

			if remaining > 0 {
				cooldown = append(
					cooldown,
					[2]int{remaining, time + n + 1},
				)
			}
		}

		if len(cooldown) > 0 && cooldown[0][1] == time+1 {
			heap.Push(h, cooldown[0][0])
			cooldown = cooldown[1:]
		}

	 }
	 return time
}
