package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2

		row := mid / cols
		col := mid % cols

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
