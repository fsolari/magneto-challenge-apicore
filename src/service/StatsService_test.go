package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/mercadolibre/magneto-challenge-apicore/src/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/src/dao"
)

func TestCalculateDNAStatsShouldReturnStats(t *testing.T) {
	var dna [2]domain.DNA
	dna[0].DNA = []string{"TTGCTA", "CTGTGC", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"}
	dna[0].Mutant = false
	dna[1].DNA = []string{"TTGCTA", "CCCCAA", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"}
	dna[1].Mutant = true

	for _, e := range dna {
		dao.InsertDNA(e)
	}

	stats, err := CalculateDNAStats()

	assert.Equal(t, 1, stats.CountHumanDna, fmt.Sprintf("CountHumanDna must be 1 %v", stats.CountHumanDna))
	assert.Equal(t, 1, stats.CountMutantDna, fmt.Sprintf("CountMutantDna must be 1 %v", stats.CountMutantDna))
	assert.Equal(t, float64(1), stats.Ratio, fmt.Sprintf("Ratio must be 1 %v", stats.Ratio))
	assert.Nil(t, err, fmt.Sprintf("err must be nil %v", err))

	for _, e := range dna {
		dao.DeleteDNA(e)
	}

}