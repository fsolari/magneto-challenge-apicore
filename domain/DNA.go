package domain

import (
	"github.com/mercadolibre/test/magneto-challenge-apicore/util"
)

type DNA struct {
	DNA    []string `json:"dna"`
	Mutant bool
}

var DNACharSet = []string{"A", "T", "C", "G"}

func IsDNAValid(dna DNA) bool {
	if len(dna.DNA) < 3 {
		return false
	}
	for _, str := range dna.DNA {
		for _, char := range str {
			if !util.Contains(DNACharSet, string(char)) {
				return false
			}
		}
	}
	return true
}