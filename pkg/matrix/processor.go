package matrix

import "math"

type Processor struct{}

// Invert inverts a matrix
// 1,2		1 4
// 4,5	->  2 5
// If the matrix cannot be inverted, return an error.
func (m *Processor) Invert(matrix Matrix) (Matrix, error) {
	if len(matrix) == 0 {
		return matrix, nil
	}

	ROW, COL := len(matrix), len(matrix[0])
	if ROW != COL {
		return Matrix{}, InverseMatrixError{ROW, COL}
	}
	invertedMatrix := make(Matrix, COL)
	for i := 0; i < COL; i++ {
		invertedMatrix[i] = make([]int, ROW)
	}

	for i := 0; i < ROW; i++ {
		for k := 0; k < COL; k++ {
			invertedMatrix[i][k] = matrix[k][i]
		}
	}
	return invertedMatrix, nil
}

// Flatten takes a matrix and flatten it into a slice of integers
// [[1,2],[3,4],[5,6]] -> [1,2,3,4,5,6]
// Empty matrix will result in an empty slice
func (m *Processor) Flatten(matrix Matrix) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var result []int
	for _, row := range matrix {
		result = append(result, row...)
	}
	return result
}

// Sum the integers in the matrix. If the matrix is empty, return 0
func (m *Processor) Sum(matrix Matrix) (int, error) {
	total := 0
	for _, row := range matrix {
		for _, col := range row {
			if col > math.MaxInt-total {
				return 0, IntOverflowError{"sum"}
			}
			total += col
		}
	}
	return total, nil
}

// Multiply gives the product of the integers in the matrix
// return 0 if matrix is empty
func (m *Processor) Multiply(matrix Matrix) (int, error) {
	total := 1
	if len(matrix) == 0 {
		return 0, nil
	}
	for _, row := range matrix {
		for _, col := range row {
			tmp := total
			total *= col
			if col != 0 && total/col != tmp {
				return 0, IntOverflowError{"multiply"}
			}
		}
	}
	return total, nil

}
