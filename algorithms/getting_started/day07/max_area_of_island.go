package day07

import "fmt"

// 给你一个大小为 m x n 的二进制矩阵 grid 。
// 岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。
// 岛屿的面积是岛上值为 1 的单元格的数目。
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
//
// 示例 1：
// 输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 输出：6
// 解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
//
// 示例 2：
// 输入：grid = [[0,0,0,0,0,0,0,0]]
// 输出：0
//
// 提示：
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// grid[i][j] 为 0 或 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-area-of-island
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// maxAreaOfIsland 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	allPoints := map[string]bool{}
	max := 0
	i, j := 0, 0
	for {
		if !allPoints[fmt.Sprintf("%d_%d", i, j)] && grid[i][j] == 1 {
			queue := [][2]int{
				{i, j},
			}
			points := map[string]bool{}
			area := 0
			for len(queue) > 0 {
				p := queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				key := fmt.Sprintf("%d_%d", p[0], p[1])
				_, ok := points[key]
				if ok {
					continue
				}
				points[fmt.Sprintf("%d_%d", p[0], p[1])] = true
				allPoints[fmt.Sprintf("%d_%d", p[0], p[1])] = true
				area++
				tmps := [][2]int{
					{p[0] - 1, p[1]},
					{p[0], p[1] - 1},
					{p[0], p[1] + 1},
					{p[0] + 1, p[1]},
				}
				for _, v := range tmps {
					if v[0] >= 0 && v[0] < len(grid) && v[1] >= 0 && v[1] < len(grid[0]) && grid[v[0]][v[1]] == 1 {
						queue = append(queue, v)
					}
				}
			}
			if area > max {
				max = area
			}
		}
		if j < len(grid[0])-1 {
			j++
		} else if i < len(grid)-1 {
			i++
			j = 0
		} else {
			break
		}
	}
	return max
}
