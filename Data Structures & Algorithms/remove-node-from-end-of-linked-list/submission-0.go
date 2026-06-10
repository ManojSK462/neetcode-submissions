/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
	dummy := &ListNode{Next:head}
	curr := head
	Length := 0

	for curr != nil{
		curr = curr.Next
		Length++
	}
// fmt.Println(Length)
	steps:=Length-n

	curr = dummy

	for i:=0;i<steps;i++{
		curr = curr.Next
	}
	// fmt.Println(curr.Val)
	curr.Next = curr.Next.Next

	return dummy.Next



}

