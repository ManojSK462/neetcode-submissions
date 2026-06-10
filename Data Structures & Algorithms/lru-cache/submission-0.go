type Node struct{
	key int
	val int
	next *Node
	prev *Node

}
type LRUCache struct {
    memory map[int]*Node
	capacity int
	l *Node
	r *Node  	
}

func Constructor(capacity int) LRUCache {
	l:=&Node{}
	r:=&Node{}
	l.next = r
	r.prev = l

    return LRUCache{
		memory: make(map[int]*Node, capacity),
		capacity: capacity,
		l: l,
		r: r,
	}
}

func (this *LRUCache) Get(key int) int {
    
		node, ok := this.memory[key]

		if !ok {return -1}

		this.memory[key].prev.next = this.memory[key].next
		this.memory[key].next.prev = this.memory[key].prev

		prev := this.r.prev
		prev.next = node
		node.prev = prev

		node.next = this.r
		this.r.prev = node		
		

		return node.val
	
}

func (this *LRUCache) Put(key int, value int) {
   
	node, ok := this.memory[key]

	if ok{
		this.memory[key].val = value

		this.memory[key].prev.next = this.memory[key].next
		this.memory[key].next.prev = this.memory[key].prev

		prev := this.r.prev
		prev.next = node
		node.prev = prev

		node.next = this.r
		this.r.prev = node	
	}else{

		if this.capacity == len(this.memory){

			dkey := this.l.next.key
			
			this.l.next = this.l.next.next
			this.l.next.prev = this.l

			delete(this.memory, dkey)
		}
		node := &Node{key:key, val: value}
		prev := this.r.prev

		prev.next = node
		node.prev = prev

		node.next = this.r
		this.r.prev = node	

		this.memory[key] = node 
	
	}


}
