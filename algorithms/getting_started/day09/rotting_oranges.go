package day09

// 在给定的网格中，每个单元格可以有以下三个值之一：
//
// 值0代表空单元格；
// 值1代表新鲜橘子；
// 值2代表腐烂的橘子。
// 每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
//
// 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回-1。
//
// 示例 1：
// 输入：[[2,1,1],[1,1,0],[0,1,1]]
// 输出：4
//
// 示例 2：
// 输入：[[2,1,1],[0,1,1],[1,0,1]]
// 输出：-1
// 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
//
// 示例 3：
// 输入：[[0,2]]
// 输出：0
// 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
// 提示：
// 1 <= grid.length <= 10
// 1 <= grid[0].length <= 10
// grid[i][j] 仅为0、1或2
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/rotting-oranges
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// orangesRotting 腐烂的橘子
func orangesRotting(grid [][]int) int {
	dr := [4]int{-1, 0, 1, 0}
	dc := [4]int{0, -1, 0, 1}
	R, C := len(grid), len(grid[0])
	var queue []int
	depth := map[int]int{}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 2 {
				code := r*C + c
				queue = append(queue, code)
				depth[code] = 0
			}
		}
	}
	ans := 0
	for len(queue) > 0 {
		code := queue[0]
		queue = queue[1:]
		r, c := code/C, code%C
		for k := 0; k < 4; k++ {
			nr := r + dr[k]
			nc := c + dc[k]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				ncode := nr*C + nc
				queue = append(queue, ncode)
				depth[ncode] = depth[code] + 1
				ans = depth[ncode]
			}
		}
	}
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				return -1
			}
		}
	}
	return ans
}
