package test

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/magneto-challenge-apicore/util"
)

func TestStringContainsIfPresentMustReturnTrue(t *testing.T) {
	var a = []string{"H","E", "L", "L", "O"}
	var c = "H"

	ok := util.Contains(a, c)

	assert.True(t, ok, fmt.Sprintf("ok must be true %v"))
}

func TestStringContainsIfNotPresentMustReturnFalse(t *testing.T) {
	var a = []string{"H","E", "L", "L", "O"}
	var c = "W"

	ok := util.Contains(a, c)

	assert.False(t, ok, fmt.Sprintf("ok must be false %v"))
}