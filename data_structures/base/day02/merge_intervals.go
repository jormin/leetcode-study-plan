package day02

import (
	"math"
	"sort"
)

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
// 示例 1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
// 示例2：
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
// 提示：
// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-intervals
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// merge 合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(
		intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		},
	)
	var ans [][]int
	for _, v := range intervals {
		if len(ans) == 0 || ans[len(ans)-1][1] < v[0] {
			ans = append(ans, v)
		} else {
			last := ans[len(ans)-1]
			ans[len(ans)-1][1] = int(math.Max(float64(last[1]), float64(v[1])))
		}
	}
	return ans
}
