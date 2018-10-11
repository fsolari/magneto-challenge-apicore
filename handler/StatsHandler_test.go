package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetDNAStatsMustReturn200AndStats(t *testing.T) {

	var currentStats = `{"count_mutant_dna":1,"count_human_dna":1,"ratio":1}`
	url := "/stats"
	r := handlerTest("GET", url, "")

	assert.Equal(t, 200, r.Code)
	assert.Equal(t, currentStats, r.Body.String())
}

