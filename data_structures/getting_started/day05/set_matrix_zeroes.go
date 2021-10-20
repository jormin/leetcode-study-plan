package day05

// 给定一个m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
// 进阶：
// 一个直观的解决方案是使用 O(mn)的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m+n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个仅使用常量空间的解决方案吗？
//
// 示例 1：
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
//
// 示例 2：
// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// 提示：
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -231 <= matrix[i][j] <= 231 - 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/set-matrix-zeroes
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// setZeroes 矩阵置零
func setZeroes(matrix [][]int) {
	row0 := false
	col0 := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			row0 = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}