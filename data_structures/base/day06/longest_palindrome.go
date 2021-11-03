package day06

// 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
// 在构造过程中，请注意区分大小写。比如"Aa"不能当做一个回文字符串。
//
// 注意:
// 假设字符串的长度不会超过 1010。
//
// 示例 1:
// 输入:
// "abccccdd"
// 输出:
// 7
// 解释:
// 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindrome
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// longestPalindrome 最长回文串
func longestPalindrome(s string) int {
	if s == "" {
		return 0
	}
	arr := make([]int, 58)
	for i := 0; i < len(s); i++ {
		arr[s[i]-'A'] ++
	}
	ans := 0
	for _, v := range arr {
		ans += v / 2 * 2
		if ans%2 == 0 && v%2 == 1 {
			ans++
		}
	}
	return ans
}
