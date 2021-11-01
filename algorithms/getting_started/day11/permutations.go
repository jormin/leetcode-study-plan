package day11

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// 示例 2：
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
//
// 示例 3：
// 输入：nums = [1]
// 输出：[[1]]
//
// 提示：
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// permute 全排列
func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	temp := make([]int, n)
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(depth int) {
		if depth == len(nums) {
			comb := make([]int, n)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		for i := 0; i < n; i++ {
			if !visited[i] {
				visited[i] = true
				temp[depth] = nums[i]
				dfs(depth + 1)
				visited[i] = false
			}
		}
	}
	dfs(0)
	return ans
}
