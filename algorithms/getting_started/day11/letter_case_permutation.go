package day11

import (
	"fmt"
	"strings"
	"unicode"
)

// 给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
//
// 示例：
// 输入：S = "a1b2"
// 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
// 输入：S = "3z4"
// 输出：["3z4", "3Z4"]
// 输入：S = "12345"
// 输出：["12345"]
//
// 提示：
// S的长度不超过12。
// S仅由数字和字母组成。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/letter-case-permutation
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// letterCasePermutation 字母大小写全排列
func letterCasePermutation(s string) []string {
	var ans []string
	for _, v := range s {
		c := string(v)
		if unicode.IsLetter(v) {
			low := strings.ToLower(c)
			up := strings.ToUpper(c)
			if len(ans) == 0 {
				ans = append(ans, low, up)
				continue
			}
			for i, s := range ans {
				ans[i] = fmt.Sprintf("%s%s", s, low)
				ans = append(ans, fmt.Sprintf("%s%s", s, up))
			}
		} else {
			if len(ans) == 0 {
				ans = append(ans, c)
				continue
			}
			for i, s := range ans {
				ans[i] = fmt.Sprintf("%s%s", s, string(v))
			}
		}
	}
	return ans
}
