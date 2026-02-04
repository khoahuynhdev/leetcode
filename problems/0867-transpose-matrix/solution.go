package solution

// https://leetcode.com/problems/transpose-matrix/solutions/4364750/golang-simple-solution/?envType=daily-question&envId=2023-12-10
func transpose(matrix [][]int) [][]int {
	newMatrixRowsCount := len(matrix[0])
	newMatrixColsCount := len(matrix)

	newMatrix := make([][]int, newMatrixRowsCount)

	for row := 0; row < newMatrixRowsCount; row++ {
		newMatrix[row] = make([]int, newMatrixColsCount)

		for col := 0; col < newMatrixColsCount; col++ {
			newMatrix[row][col] = matrix[col][row]
		}
	}

	return newMatrix
}
