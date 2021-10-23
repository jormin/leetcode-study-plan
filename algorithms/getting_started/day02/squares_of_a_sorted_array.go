package day02

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
// 示例 1：
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
//
// 示例 2：
// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
//
// 提示：
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 已按 非递减顺序 排序
//
// 进阶：
// 请你设计时间复杂度为 O(n) 的算法解决本问题
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// sortedSquares 有序数组的平方
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		v1 := nums[i] * nums[i]
		v2 := nums[j] * nums[j]
		if v1 > v2 {
			ans[pos] = v1
			i++
		} else {
			ans[pos] = v2
			j--
		}
	}
	return ans
}
