package day03

// 给你一个正整数n ，生成一个包含 1 到n2所有元素，且元素按顺时针顺序螺旋排列的n x n 正方形矩阵 matrix 。
//
// 示例 1：
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 示例 2：
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
// 1 <= n <= 20
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/spiral-matrix-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// generateMatrix 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, dirIndex := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		ans[row][col] = i
		dir := dirs[dirIndex]
		// 如果下一个位置超出边界或者已经有值，则顺时针进行旋转
		if r, c := row+dir[0], col+dir[1]; r < 0 || r >= n || c < 0 || c >= n || ans[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = dirs[dirIndex]
		}
		row += dir[0]
		col += dir[1]
	}
	return ans
}
