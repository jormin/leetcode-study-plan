package day03

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
// 示例 1:
// 输入: rowIndex = 3
// 输出: [1,3,3,1]
//
// 示例 2:
// 输入: rowIndex = 0
// 输出: [1]
//
// 示例 3:
// 输入: rowIndex = 1
// 输出: [1,1]
//
// 提示:
// 0 <= rowIndex <= 33
//
// 进阶：
// 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/pascals-triangle-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// getRow 杨辉三角 II
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}
	return ans
}
