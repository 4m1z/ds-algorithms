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
func bfs(matrix *[][]int, initialRow, initialCol int) {
	queue := [][]int{{initialRow, initialCol}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		row := current[0]
		col := current[1]

		for _, direction := range directions {
			tempRow := row + direction[0]
			tempCol := col + direction[1]
			if tempRow < len(*matrix) && tempRow >= 0 && tempCol < len((*matrix)[0]) && tempCol >= 0 {
				if (*matrix)[tempRow][tempCol] == 1 {
					(*matrix)[tempRow][tempCol] = 0
					queue = append(queue, []int{tempRow, tempCol})
				}
			}
		}
	}
}

func countIslandsBfs(matrix [][]int) (t int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				t++
				bfs(&matrix, i, j)
			}
		}
	}

	return t
}
