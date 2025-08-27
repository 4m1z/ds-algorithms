package main

func dfs40(m [][]int, i, j int) {
	for _, dir := range directions {
		ti := i + dir[0]
		tj := j + dir[1]

		if m[ti][tj] == 1 {
			m[ti][tj] = 0
			dfs40(m, ti, tj)
		}
	}
}

func countIslandsDfs(m [][]int) (t int) {
	for i, row := range m {
		for j, col := range row {
			if col == 1 {
				t++
				dfs40(m, i, j)
			}
		}
	}

	return t
}

// / bfs
func bfs(m [][]int, i, j int) {
	queue := [][]int{{i, j}}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			ti := c[0] + dir[0]
			tj := c[1] + dir[1]

			if m[ti][tj] == 1 {
				m[ti][tj] = 0
				queue = append(queue, []int{ti, tj})
			}

		}
	}
}

func countIslandsBfs(matrix [][]int) (t int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				t++
				bfs(matrix, i, j)
			}
		}
	}

	return t
}
