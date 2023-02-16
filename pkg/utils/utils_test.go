package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntArrToStr(t *testing.T) {
	s := IntArrToStr([]int{1, 2, 3})
	expectedS := "1,2,3"
	assert.Equal(t, expectedS, s)

	s = IntArrToStr([]int{})
	expectedS = ""
	assert.Equal(t, expectedS, s)
}
