package test

import (
	"testing"

	"github.com/mercadolibre/test/magneto-challenge-apicore/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestIsDNAValidInvalidLengthMustReturnFalse(t *testing.T) {
	var dna domain.DNA
	dna.DNA = append(dna.DNA, "ATGCGA")

	ok := domain.IsDNAValid(dna)

	assert.False(t, ok, fmt.Sprintf("ok must be false %v"))
}

func TestIsDNAValidInvalidCharMustReturnFalse(t *testing.T) {
	var dna domain.DNA
	dna.DNA = append(dna.DNA, "ATGCGA", "CARRGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG")

	ok := domain.IsDNAValid(dna)

	assert.False(t, ok, fmt.Sprintf("ok must be false %v"))

}

func TestIsDNAValidIsValidMustReturnTrue(t *testing.T) {
	var dna domain.DNA
	dna.DNA = append(dna.DNA, "ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG")

	ok := domain.IsDNAValid(dna)

	assert.True(t, ok, fmt.Sprintf("ok must be true %v"))


}