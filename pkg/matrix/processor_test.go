package matrix

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHelper_Invert(t *testing.T) {
	helper := Processor{}

	// normal case
	m := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mInverse, err := helper.Invert(m)
	expectedResult := Matrix{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	assert.Nil(t, err)
	assert.Equal(t, mInverse, expectedResult)

	// no inverse possible
	m = Matrix{{1, 2, 3}, {4, 5, 6}}
	_, err = helper.Invert(m)
	assert.IsTypef(t, InverseMatrixError{}, err, "expected InverseMatrixError not found")
}

func TestHelper_Flatten(t *testing.T) {
	helper := Processor{}

	// normal case
	m := Matrix{{1, 2, 3}, {4, 5, 6}}
	mFlatten := helper.Flatten(m)
	expectedResult := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, mFlatten, expectedResult, "unexpected flatten result")

	// empty matrix
	m = Matrix{}
	mFlatten = helper.Flatten(m)
	expectedResult = []int{}
	assert.Equal(t, mFlatten, expectedResult, "unexpected flatten")
}

func TestHelper_Sum(t *testing.T) {
	helper := Processor{}

	// normal case
	m := Matrix{{1, 2, 3}, {4, 5, 6}}
	total, err := helper.Sum(m)
	expectedResult := 21
	assert.Nil(t, err)
	assert.Equal(t, total, expectedResult)

	// overflow case
	m = Matrix{{1, 10, 1000}, {math.MaxInt, 100, math.MaxInt}}
	_, err = helper.Sum(m)
	assert.IsTypef(t, IntOverflowError{}, err, "expcted IntOverflowError not found")

}

func TestHelper_Multiply(t *testing.T) {
	helper := Processor{}

	// normal case
	m := Matrix{{1, 2}, {4, 5}}
	total, err := helper.Multiply(m)
	expectedResult := 40
	assert.Nil(t, err)
	assert.Equal(t, total, expectedResult)

	// overflow case
	m = Matrix{{1, 10}, {math.MaxInt, math.MaxInt}}
	_, err = helper.Multiply(m)
	assert.IsTypef(t, IntOverflowError{}, err, "expcted IntOverflowError not found")

}
