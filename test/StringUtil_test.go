package test

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/magneto-challenge-apicore/util"
)

func TestStringContainsIfPresentMustReturnTrue(t *testing.T) {
	var a = []string{"H", "E", "L", "L", "O"}
	var c = "H"

	ok := util.StringArrayContains(a, c)

	assert.True(t, ok, fmt.Sprintf("ok must be true %v", ok))
}

func TestStringContainsIfNotPresentMustReturnFalse(t *testing.T) {
	var a = []string{"H", "E", "L", "L", "O"}
	var c = "W"

	ok := util.StringArrayContains(a, c)

	assert.False(t, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestStringToRuneArrayMustReturnRuneArray(t *testing.T) {
	var s = "CAGTGC"
	var r = []rune{67, 65, 71, 84, 71, 67}

	ok := util.StringToRuneArray(s)

	assert.Equal(t, r, ok, fmt.Sprintf("ok must be in rune array type %v", ok))
}

func TestGenerateMatrixFromStringArrayMustReturnRuneMatrix(t *testing.T) {
	var arr = []string{"TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	var matrix = [][]rune{{84, 84, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 65, 71, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	ok := util.GenerateMatrixFromStringArray(arr)

	assert.Equal(t, matrix, arr, fmt.Sprintf("ok must be matrix rune type %v", ok))
}