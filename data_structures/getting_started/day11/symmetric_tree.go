package day11

// 给定一个二叉树，检查它是否是镜像对称的。
//
// 例如，二叉树[1,2,2,3,4,4,3] 是对称的。
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
// 但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
// 进阶：
//
// 你可以运用递归和迭代两种方法解决这个问题吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/symmetric-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// isSymmetric 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return _isSymmetric(root.Left, root.Right)
}

// _isSymmetric 检查是否镜像对称
func _isSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && _isSymmetric(left.Left, right.Right) && _isSymmetric(left.Right, right.Left)
}
