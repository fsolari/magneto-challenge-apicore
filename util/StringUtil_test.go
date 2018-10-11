package util

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestStringContainsIfPresentMustReturnTrue(t *testing.T) {
	var a = []string{"H", "E", "L", "L", "O"}
	var c = "H"

	ok := StringArrayContains(a, c)

	assert.True(t, ok, fmt.Sprintf("ok must be true %v", ok))
}

func TestStringContainsIfNotPresentMustReturnFalse(t *testing.T) {
	var a = []string{"H", "E", "L", "L", "O"}
	var c = "W"

	ok := StringArrayContains(a, c)

	assert.False(t, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestStringToRuneArrayMustReturnRuneArray(t *testing.T) {
	var s = "CAGTGC"
	var r = []rune{67, 65, 71, 84, 71, 67}

	ok := stringToRuneArray(s)

	assert.Equal(t, r, ok, fmt.Sprintf("ok must be in rune array type %v", ok))
}

func TestJoinStringArrayMustReturnStringJoined(t *testing.T) {

	var a = []string{"H", "E", "L", "L", "O"}
	var c = "H E L L O"

	ok := JoinStringArray(a)
	assert.Equal(t, ok, c, fmt.Sprintf("ok must equal c %v %v", ok, c))
}
