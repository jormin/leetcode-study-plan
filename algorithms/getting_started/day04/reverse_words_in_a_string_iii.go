package day04

import (
	"strings"
)

// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
// 示例：
// 输入："Let's take LeetCode contest"
// 输出："s'teL ekat edoCteeL tsetnoc"
//
// 提示：
// 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// reverseWords 反转字符串中的单词 III
func reverseWords(s string) string {
	arr := strings.Split(s, " ")

	for i, v := range arr {
		b := []byte(v)
		n := len(b)
		left, right := 0, n-1
		for left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
		arr[i] = string(b)
	}
	return strings.Join(arr, " ")
}
