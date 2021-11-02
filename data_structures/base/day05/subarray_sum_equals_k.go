package day05

// 给你一个整数数组 nums 和一个整数k ，请你统计并返回该数组中和为k的连续子数组的个数。
//
// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2
//
// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2
//
// 提示：
// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// subarraySum 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
