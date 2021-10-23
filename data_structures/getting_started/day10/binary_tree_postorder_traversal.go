package day10

// 给定一个二叉树，返回它的 后序遍历。
//
// 示例:
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
// 输出: [3,2,1]
// 进阶:递归算法很简单，你可以通过迭代算法完成吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// postorderTraversal 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var nodes []*TreeNode
	var prev *TreeNode
	for root != nil || len(nodes) > 0 {
		for root != nil {
			nodes = append(nodes, root)
			root = root.Left
		}
		root = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			nodes = append(nodes, root)
			root = root.Right
		}
	}
	return ans
}
