package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/dao"
)

func TestGetDNAStatsMustReturn200AndStats(t *testing.T) {

	var dna [2]domain.DNA
	dna[0].DNA = []string{"TTGCTA", "CTGTGC", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"}
	dna[0].Mutant = false
	dna[1].DNA = []string{"TTGCTA", "CCCCAA", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"}
	dna[1].Mutant = true

	for _, e := range dna {
		dao.InsertDNA(e)
	}

	var currentStats = `{"count_mutant_dna":1,"count_human_dna":1,"ratio":1}`
	url := "/stats"
	r := handlerTest("GET", url, "")

	assert.Equal(t, 200, r.Code)
	assert.Equal(t, currentStats, r.Body.String())

	for _, e := range dna {
		dao.DeleteDNA(e)
	}
}

