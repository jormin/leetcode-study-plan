package day10

// 给定一个二叉树的根节点 root ，返回它的 中序遍历。
//
// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
// 示例 2：
// 输入：root = []
// 输出：[]
//
// 示例 3：
// 输入：root = [1]
// 输出：[1]
//
// 示例 4：
// 输入：root = [1,2]
// 输出：[2,1]
//
// 示例 5：
// 输入：root = [1,null,2]
// 输出：[1,2]
//
// 提示：
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶:递归算法很简单，你可以通过迭代算法完成吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// inorderTraversal 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var nodes []*TreeNode
	for root != nil || len(nodes) > 0 {
		for root != nil {
			nodes = append(nodes, root)
			root = root.Left
		}
		root = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
