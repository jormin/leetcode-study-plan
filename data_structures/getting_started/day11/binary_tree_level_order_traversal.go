package day11

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
// 示例：
// 二叉树：[3,9,20,null,null,15,7],
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 返回其层序遍历结果：
// [
//  [3],
//  [9,20],
//  [15,7]
// ]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	p := []*TreeNode{root}
	for i := 0; len(p) > 0; i++ {
		ans = append(ans, []int{})
		var q []*TreeNode
		for j := 0; j < len(p); j++ {
			node := p[j]
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
	}
	return ans
}
