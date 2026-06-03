type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.size {
		return -1
	}

	curr := ll.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.val
}

func (ll *LinkedList) InsertHead(val int) {
	node := &Node{val: val}

	if ll.size == 0 {
		ll.head = node
		ll.tail = node
	} else {
		node.next = ll.head
		ll.head = node
	}

	ll.size++
}

func (ll *LinkedList) InsertTail(val int) {
	node := &Node{val: val}

	if ll.size == 0 {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	ll.size++
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || index >= ll.size {
		return false
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.size--

		if ll.size == 0 {
			ll.tail = nil
		}

		return true
	}

	prev := ll.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	removed := prev.next
	prev.next = removed.next

	if removed == ll.tail {
		ll.tail = prev
	}

	ll.size--
	return true
}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, 0, ll.size)

	curr := ll.head
	for curr != nil {
		values = append(values, curr.val)
		curr = curr.next
	}

	return values
}