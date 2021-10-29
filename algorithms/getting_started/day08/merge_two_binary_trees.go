package day08

// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。
// 示例1:
// 输入:
//	Tree 1                     Tree 2
//          1                         2
//         / \                       / \
//        3   2                     1   3
//       /                           \   \
//      5                             4   7
// 输出:
// 合并后的树:
//	     3
//	    / \
//	   4   5
//	  / \   \
//	 5   4   7
// 注意:合并必须从两个树的根节点开始。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-binary-trees
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mergeTrees 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
