package day06

// 给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
// 示例 1：
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
//
// 示例 2：
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false
//
// 提示：
// 1 <= s1.length, s2.length <= 104
// s1 和 s2 仅包含小写字母
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutation-in-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// checkInclusion 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	b1, b2 := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		b1[s1[i]-'a']++
		b2[s2[i]-'a']++
	}
	if b1 == b2 {
		return true
	}
	for i := n; i < m; i++ {
		b2[s2[i-n]-'a']--
		b2[s2[i]-'a']++
		if b1 == b2 {
			return true
		}
	}
	return false
}
