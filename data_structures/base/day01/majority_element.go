package day01

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
// 示例1：
// 输入：[3,2,3]
// 输出：3
//
// 示例2：
// 输入：[2,2,1,1,1,2,2]
// 输出：2
//
// 进阶：
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/majority-element
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// majorityElement 多数元素
func majorityElement(nums []int) int {
	ans, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			ans = v
		}
		if ans == v {
			count++
		} else {
			count--
		}
	}
	return ans
}
