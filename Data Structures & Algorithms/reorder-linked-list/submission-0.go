/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    
	if head == nil || head.Next ==nil {
		return 
	}



	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
	}

	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for  second!= nil{
		next := second.Next
		second.Next = prev
		prev = second
		second = next
	}

	first := head

	second = prev

	for second != nil{
		n1 := first.Next
		n2 := second.Next

		first.Next = second
		second.Next = n1

		first = n1
		second = n2


	}


}


