package domain

import (
	"github.com/mercadolibre/magneto-challenge-apicore/util"
)

var DNACharSet = []string{"A", "T", "C", "G"}

type DNA struct {
	DNA    []string `json:"dna;sql:"not null;unique"`
	Mutant bool
}

func IsDNAValid(dna DNA) bool {
	var rows, cols int = len(dna.DNA), len(dna.DNA[0])

	if rows < 3 || rows != cols {
		return false
	}

	for _, str := range dna.DNA {
		for _, char := range str {
			if !util.StringArrayContains(DNACharSet, string(char)) {
				return false
			}
		}
	}
	return true
}