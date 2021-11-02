package day12

import "math"

// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
// 示例 1：
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//   2
//  3 4
// 6 5 7
// 4 1 8 3
// 自顶向下的最小路径和为11（即，2+3+5+1= 11）。
//
// 示例 2：
// 输入：triangle = [[-10]]
// 输出：-10
//
// 提示：
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -104 <= triangle[i][j] <= 104
//
// 进阶：
// 你可以只使用 O(n)的额外空间（n 为三角形的总行数）来解决这个问题吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/triangle
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// minimumTotal 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n)
	f[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			f[j] = minStep(f[j-1], f[j]) + triangle[i][j]
		}
		f[0] += triangle[i][0]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = minStep(ans, f[i])
	}
	return ans
}

func minStep(x, y int) int {
	if x < y {
		return x
	}
	return y
}
