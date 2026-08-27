# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def valid(n, l, r):
            if not n: return True
            if not l<n.val<r:
                return False
            
            return valid(n.left, l, n.val) and valid(n.right, n.val, r)
            
        return valid(root, float('-inf'), float('inf'))


        