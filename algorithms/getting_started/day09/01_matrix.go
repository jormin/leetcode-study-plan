package day09

import (
	"math"
)

// 给定一个由 0 和 1 组成的矩阵 mat，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
//
// 两个相邻元素间的距离为 1 。
//
// 示例 1：
// 输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：[[0,0,0],[0,1,0],[0,0,0]]
//
// 示例 2：
// 输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
// 输出：[[0,0,0],[0,1,0],[1,2,1]]
//
// 提示：
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 104
// 1 <= m * n <= 104
// mat[i][j] is either 0 or 1.
// mat 中至少有一个 0
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/01-matrix
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// updateMatrix 01 矩阵
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				ans[i][j] = 0
			} else {
				ans[i][j] = math.MaxInt32
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				ans[i][j] = int(math.Min(float64(ans[i][j]), float64(ans[i-1][j]+1)))
			}
			if j-1 >= 0 {
				ans[i][j] = int(math.Min(float64(ans[i][j]), float64(ans[i][j-1]+1)))
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				ans[i][j] = int(math.Min(float64(ans[i][j]), float64(ans[i+1][j]+1)))
			}
			if j+1 < n {
				ans[i][j] = int(math.Min(float64(ans[i][j]), float64(ans[i][j+1]+1)))
			}
		}
	}
	return ans
}
