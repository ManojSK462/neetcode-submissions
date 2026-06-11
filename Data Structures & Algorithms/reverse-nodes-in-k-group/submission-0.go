/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    
    dummy := &ListNode{Next:head}

    curr := dummy

    for curr!=nil {
		wPrev:= curr
		wStart := curr.Next
		wEnd:= getKth(curr, k)

		if wEnd==nil{
			break
		}
		prev:= wEnd.Next
		curr = wStart

		for prev != wEnd {
			next:= curr.Next

			curr.Next = prev

			prev = curr

			curr = next
			}
			wPrev.Next = wEnd
			curr = wStart
		}
    return dummy.Next

}
func getKth(curr *ListNode, k int) *ListNode{
    for k>0 && curr!=nil{
        curr = curr.Next
        k--
    }

    return curr
}