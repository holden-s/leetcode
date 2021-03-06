给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

示例 1：

输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32
示例 2：

输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23
 

提示：

树中节点数目在范围 [1, 2 * 104] 内
1 <= Node.val <= 105
1 <= low <= high <= 105
所有 Node.val 互不相同

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var rangeSumBSTHelper func(*TreeNode)
	rangeSumBSTHelper = func(r *TreeNode) {
		if r == nil {
			return
		}	
		if r.Val < low {
			rangeSumBSTHelper(r.Right)
			return
		}
		rangeSumBSTHelper(r.Left)
		if r.Val >= low && r.Val <= high {
			sum += r.Val
			rangeSumBSTHelper(r.Right)
		}
	}
	rangeSumBSTHelper(root)

	return sum
}
