package day06

import (
	"fmt"
	"strconv"
	"strings"
)

// 给定两个字符串形式的非负整数num1 和num2，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger），也不能直接将输入的字符串转换为整数形式。
//
// 示例 1：
// 输入：num1 = "11", num2 = "123"
// 输出："134"
//
// 示例 2：
// 输入：num1 = "456", num2 = "77"
// 输出："533"
//
// 示例 3：
// 输入：num1 = "0", num2 = "0"
// 输出："0"
//
// 提示：
// 1 <= num1.length, num2.length <= 104
// num1 和num2 都只包含数字0-9
// num1 和num2 都不包含任何前导零
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-strings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// addStrings 字符串相加
func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	length := m
	if n > m {
		length = n
	}
	arr1 := make([]int, length)
	arr2 := make([]int, length)
	for i := m - 1; i >= 0; i-- {
		arr1[m-1-i], _ = strconv.Atoi(string(num1[i]))
	}
	for i := n - 1; i >= 0; i-- {
		arr2[n-1-i], _ = strconv.Atoi(string(num2[i]))
	}
	ans := make([]int, length+1)
	f := false
	for i := 0; i < length; i++ {
		tmp := arr1[i] + arr2[i]
		if f {
			tmp += 1
		}
		if tmp >= 10 {
			f = true
			tmp -= 10
		} else {
			f = false
		}
		ans[i] = tmp
	}
	if f {
		ans[len(ans)-1] = 1
	}
	res := ""
	for i := len(ans) - 1; i >= 0; i-- {
		res += fmt.Sprintf("%d", ans[i])
	}
	res = strings.TrimLeft(res, "0")
	if res == "" {
		res = "0"
	}
	return res
}
