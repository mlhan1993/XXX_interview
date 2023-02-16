package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// InverseMatrixError indicates matrix cannot be inverted
type InverseMatrixError struct {
	m int
	n int
}

func (e InverseMatrixError) Error() string {
	return fmt.Sprintf("matrix of size %d x %d does not have inverse", e.m, e.n)
}

// InvalidMatrixError indicates the input matrix is an invalid matrix
type InvalidMatrixError struct{}

func (e InvalidMatrixError) Error() string {
	return fmt.Sprintf("matrix is not a valid matrix")
}

// IntOverflowError indicates doing this operation will result in an int overflow
type IntOverflowError struct {
	operation string
}

func (e IntOverflowError) Error() string {
	return fmt.Sprintf("doing %s operation will result in integer overflow", e.operation)
}

type Matrix [][]int

// ToString return the string representation of the matrix
// 1,2,3
// 3,4,5
// 6,7,8
func (m Matrix) ToString() string {
	var rows []string
	for _, row := range m {
		var strArr []string
		for _, col := range row {
			num := strconv.Itoa(col)
			strArr = append(strArr, num)
		}
		rows = append(rows, strings.Join(strArr, ","))
	}
	return strings.Join(rows, "\n")
}

// Validate returns an error when the matrix is not a valid matrix (ie. [[1],[1,2,3]]
func (m Matrix) Validate() error {
	myLen := 0
	for _, row := range m {
		if myLen == 0 {
			myLen = len(row)
		} else {
			if myLen != len(row) {
				return InvalidMatrixError{}
			}
		}
	}
	return nil
}
