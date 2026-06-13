/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := [][2]*TreeNode{{p, q}}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		first, second := pair[0], pair[1]

		if first == nil && second == nil {
			continue
		}

		if first == nil || second == nil || first.Val != second.Val {
			return false
		}

		queue = append(queue,
			[2]*TreeNode{first.Left, second.Left},
			[2]*TreeNode{first.Right, second.Right},
		)
	}

	return true
}