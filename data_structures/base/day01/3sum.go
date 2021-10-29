package day01

import (
	"sort"
)

// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
// 示例 2：
// 输入：nums = []
// 输出：[]
//
// 示例 3：
// 输入：nums = [0]
// 输出：[]
//
// 提示：
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// threeSum 三数之和
func threeSum(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}
