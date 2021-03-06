package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/dao"
)

func TestGetDNATestMustReturn200IfMutant(t *testing.T){

	var dna domain.DNA
	dna.DNA = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	body:= `{"dna":["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}`
	url := "/mutant"
	r := handlerTest("POST", url, body)

	assert.Equal(t, 200, r.Code)

	dao.DeleteDNA(dna)

}


func TestGetDNATestMustReturn403IfHuman(t *testing.T){

	var dna domain.DNA
	dna.DNA = []string{"TTGCTA", "CTGTGC", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"}

	body:= `{"dna":["TTGCTA", "CTGTGC", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"]}`
	url := "/mutant"
	r := handlerTest("POST", url, body)

	assert.Equal(t, 403, r.Code)

	dao.DeleteDNA(dna)
}
