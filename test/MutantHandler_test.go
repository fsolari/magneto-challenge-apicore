package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"net/http"
)


/*TODO: FINISH TESTS*/

func TestIsMutantForMutantMustReturn200(t *testing.T) {

	var body = `{
			"dna": "ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"
	}`

	var url = "/mutant"

	r := executeRequest("POST", url, body)

	assert.Equal(t, http.StatusOK, r.Code, fmt.Sprintf("http status should be 200 %v"))
	assert.NotNil(t, r.Body, fmt.Sprintf("response body must be not nil %v"))
}


func TestIsMutantForHumanMustReturn400(t *testing.T) {

	var body = `{
			"dna": "ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"
	}`

	var url = "/mutant"

	r := executeRequest("POST", url, body)

	assert.Equal(t, http.StatusBadRequest, r.Code, fmt.Sprintf("http status should be 200 %v"))
	assert.NotNil(t, r.Body, fmt.Sprintf("response body must be not nil %v"))
}

func TestIsMutantForHumanDNAInvalidMustReturn400(t *testing.T) {

	var body = `{
			"dna": "ATGC", "TCACTG"
	}`

	var url = "/mutant"

	r := executeRequest("POST", url, body)

	assert.Equal(t, http.StatusBadRequest, r.Code, fmt.Sprintf("http status should be 200 %v"))
	assert.NotNil(t, r.Body, fmt.Sprintf("response body must be not nil %v"))
}
