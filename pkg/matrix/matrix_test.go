package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_Validate(t *testing.T) {
	invalid := Matrix{[]int{1, 2, 3}, []int{1, 2}, []int{1, 2, 3}}

	err := invalid.Validate()
	assert.IsTypef(t, InvalidMatrixError{}, err, "invalid matrix expect to fail Validate() with InvalidMatrixError")
}

func TestMatrix_ToString(t *testing.T) {
	m := Matrix{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	str := m.ToString()
	expectedStr := "1,2,3\n4,5,6\n7,8,9"
	assert.Equalf(t, expectedStr, str, "unexpected result from ToString")
}
