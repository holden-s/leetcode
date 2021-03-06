// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
// 
// 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同
// 
//  
// 
// 示例 1：
// 
// 
// 输入：root = [4,2,6,1,3]
// 输出：1
// 示例 2：
// 
// 
// 输入：root = [1,0,48,null,null,12,49]
// 输出：1
//  
// 
// 提示：
// 
// 树中节点数目在范围 [2, 100] 内
// 0 <= Node.val <= 105

func minDiffInBST(root *TreeNode) int {
	pre, ans := -1, math.MaxInt64
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val - pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)

	return ans
}
