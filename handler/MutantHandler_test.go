package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetDNATestMustReturn200IfMutant(t *testing.T){

	body:= `{"dna":["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}`
	url := "/mutant"
	r := handlerTest("POST", url, body)

	assert.Equal(t, 200, r.Code)
}


func TestGetDNATestMustReturn403IfHuman(t *testing.T){

	body:= `{"dna":["TTGCTA", "CTGTGC", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"]}`
	url := "/mutant"
	r := handlerTest("POST", url, body)

	assert.Equal(t, 403, r.Code)
}
