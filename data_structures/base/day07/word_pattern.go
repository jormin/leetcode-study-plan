package day07

import "strings"

// 给定一种规律 pattern和一个字符串str，判断 str 是否遵循相同的规律。
//
// 这里的遵循指完全匹配，例如，pattern里的每个字母和字符串str中的每个非空单词之间存在着双向连接的对应规律。
//
// 示例1:
// 输入: pattern = "abba", str = "dog cat cat dog"
// 输出: true
//
// 示例 2:
// 输入:pattern = "abba", str = "dog cat cat fish"
// 输出: false
//
// 示例 3:
// 输入: pattern = "aaaa", str = "dog cat cat dog"
// 输出: false
//
// 示例4:
// 输入: pattern = "abba", str = "dog dog dog dog"
// 输出: false
// 说明:
// 你可以假设pattern只包含小写字母，str包含了由单个空格分隔的小写字母。 
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-pattern
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// wordPattern 单词规律
func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(pattern) != len(arr) {
		return false
	}
	m1 := map[string]byte{}
	m2 := map[byte]string{}
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		str := arr[i]
		if m1[str] > 0 && m1[str] != ch || m2[ch] != "" && m2[ch] != str {
			return false
		}
		m1[str] = ch
		m2[ch] = str
	}
	return true
}
