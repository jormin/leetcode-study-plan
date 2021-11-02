package day05

// 给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。
//
// 示例:
// 输入: [1,2,3,4]
// 输出: [24,12,8,6]
//
// 提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
//
// 说明: 请不要使用除法，且在O(n) 时间复杂度内完成此题。
//
// 进阶：
// 你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/product-of-array-except-self
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// productExceptSelf 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
