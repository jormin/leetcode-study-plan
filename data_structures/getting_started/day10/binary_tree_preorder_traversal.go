package day10

// 给你二叉树的根节点 root ，返回它节点值的前序遍历。
//
// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
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
// 输出：[1,2]
//
// 示例 5：
// 输入：root = [1,null,2]
// 输出：[1,2]
//
// 提示：
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversal 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var ans []int
	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
	}
	return ans
}
