package day11

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。
//
// 示例 1：
// 输入：n = 4, k = 2
// 输出：
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
//
// 示例 2：
// 输入：n = 1, k = 1
// 输出：[[1]]
//
// 提示：
// 1 <= n <= 20
// 1 <= k <= n
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combinations
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// combine 组合
func combine(n int, k int) [][]int {
	var ans [][]int
	var temp []int
	var dfs func(int)
	dfs = func(cur int) {
		if len(temp)+n-cur+1 < k {
			return
		}
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return ans
}
