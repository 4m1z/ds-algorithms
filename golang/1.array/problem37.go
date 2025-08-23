package main

// [[1,2,3],[4,5,6],[7,8,9]]
func SpiralOrder(matrix [][]int) (ret []int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}

	rows, cols := len(matrix), len(matrix[0])
	top, left := 0, 0
	bottom, right := rows-1, cols-1

	for left <= right && top <= bottom {
		// left to right
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top++

		// top to bottom
		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--

		// right to left
		if top <= bottom {
			for i := right; i >= left; i-- {
				ret = append(ret, matrix[bottom][i])
			}
			bottom--
		}

		// bottom to top
		if left <= right {
			for i := bottom; i >= top; i-- {
				ret = append(ret, matrix[i][left])
			}
			left++
		}
	}

	return ret
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SpiralOrder2(matrix [][]int) (ret []int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}

	for len(matrix) > 0 {
		// append first row to ret and remove it from matrix
		ret = append(ret, matrix[0]...)
		matrix = matrix[1:]

		// append last element of each inner array and remove it from that array
		if len(matrix) > 0 && len(matrix[0]) > 0 {
			for i, row := range matrix {
				col := row[len(row)-1]
				ret = append(ret, col)
				matrix[i] = row[:len(matrix[i])-1]
			}
		}

		// append reverse of the last row
		reverse(matrix[len(matrix)-1])
		ret = append(ret, matrix[len(matrix)-1]...)
		matrix = matrix[:len(matrix)-1]

		// append first element of every inner array and remove it from each one of them => bottom to top
		if len(matrix) > 0 && len(matrix[0]) > 0 {
			for i := len(matrix) - 1; i >= 0; i-- {
				col := matrix[i][0]
				ret = append(ret, col)
				matrix[i] = matrix[i][1:]
			}
		}

	}

	return ret
}
