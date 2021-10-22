package day01

import "fmt"

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
//
// 示例2:
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
//
// 示例 3:
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4
//
// 示例 4:
// 输入: nums = [1,3,5,6], target = 0
// 输出: 0
//
// 示例 5:
// 输入: nums = [1], target = 0
// 输出: 0
//
// 提示:
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为无重复元素的升序排列数组
// -104 <= target <= 104
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/search-insert-position
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// searchInsert 搜索插入位置
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	pos := len(nums)
	for low <= high {
		mid := (high-low)/2 + low
		fmt.Println(mid)
		if nums[mid] >= target {
			pos = mid
			high = mid - 1
			continue
		}
		low = mid + 1
	}
	return pos
}
