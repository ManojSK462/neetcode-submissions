/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
	if head==nil{
		return nil
	}

	translator := make(map[*Node]*Node)

	curr := head

	for curr!=nil{
		translator[curr] = &Node{Val:curr.Val}
		curr = curr.Next
	}

	curr = head

	for curr!=nil{
		translator[curr].Next = translator[curr.Next]
		translator[curr].Random = translator[curr.Random]

		curr = curr.Next
	}

	return translator[head]




}

