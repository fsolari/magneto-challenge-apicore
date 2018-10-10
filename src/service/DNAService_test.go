package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/mercadolibre/magneto-challenge-apicore/src/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/src/dao"
)

func TestIsMutantMustReturnTrueIfDNAMutant(t *testing.T) {
	var matrix = [][]rune{{65, 84, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 65, 71, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	ok := isMutant(matrix, len(matrix))

	assert.Equal(t, true, ok, fmt.Sprintf("ok must be true %v", ok))

}

func TestIsMutantMustReturnFalseIfDNAHuman(t *testing.T) {

	var matrix = [][]rune{{84, 65, 71, 67, 71, 65}, {67, 65, 84, 84, 71, 67}, {84, 84, 65, 84, 65, 84}, {65, 71, 65, 67, 71, 84}, {71, 67, 67, 65, 84, 65}, {84, 67, 65, 67, 84, 71}}

	ok := isMutant(matrix, len(matrix))

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestDNATestMustReturnFalseIfDNAIsPresent(t *testing.T) {
	var dna domain.DNA
	dna.DNA = []string{"TTGCTA", "CTGTGC", "TTATGT", "AGTAGG", "ACCCTA", "TCACTG"}
	dna.Mutant = false

	dao.InsertDNA(dna)

	ok, err := DNATest(dna)

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
	assert.Nil(t, err, fmt.Sprintf("err must be nil %v", err))

	dao.DeleteDNA(dna)

}

func TestLoopHorizontallyMustSendToChannelSequenceBeforeDone(t *testing.T) {

	var matrix = [][]rune{{65, 65, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 67, 71, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	cols := len(matrix)

	sequence := make(chan bool)
	done := make(chan bool)

	go loopHorizontally(matrix, cols, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, true, ok, fmt.Sprintf("ok must be true %v", ok))

}

func TestLoopHorizontallyIfNoSequenceMustSendToDoneChannel(t *testing.T) {

	var matrix = [][]rune{{84, 84, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 65, 71, 71}, {67, 67, 71, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	cols := len(matrix)

	sequence := make(chan bool)
	done := make(chan bool)

	go loopHorizontally(matrix, cols, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestLoopVerticallyMustSendToChannelSequenceBeforeDone(t *testing.T) {

	var matrix = [][]rune{{65, 65, 71, 67, 71, 65}, {65, 65, 71, 84, 71, 67}, {65, 84, 65, 84, 71, 84}, {65, 71, 65, 67, 71, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	cols := len(matrix)

	sequence := make(chan bool)
	done := make(chan bool)

	go loopVertically(matrix, cols, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, true, ok, fmt.Sprintf("ok must be true %v", ok))
}

func TestLoopVerticallyIfNoSequenceMustSendToDoneChannel(t *testing.T) {
	var matrix = [][]rune{{71, 65, 71, 67, 71, 65}, {65, 65, 71, 84, 71, 67}, {65, 84, 65, 84, 71, 84}, {65, 71, 65, 67, 65, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	cols := len(matrix)

	sequence := make(chan bool)
	done := make(chan bool)

	go loopVertically(matrix, cols, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestLoopDiagonallyMustSendToChannelSequenceBeforeDone(t *testing.T) {
	var matrix = [][]rune{{71, 65, 71, 67, 71, 65}, {71, 71, 71, 84, 71, 67}, {65, 84, 71, 84, 71, 84}, {65, 71, 65, 71, 65, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	sequence := make(chan bool)
	done := make(chan bool)

	go loopDiagonally(matrix, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, true, ok, fmt.Sprintf("ok must be true %v", true))

}

func TestLoopDiagonallyIfNoSequenceMustSendToDoneChannel(t *testing.T) {
	var matrix = [][]rune{{71, 65, 71, 67, 71, 65}, {65, 65, 71, 84, 71, 67}, {65, 84, 65, 84, 71, 84}, {65, 71, 65, 67, 65, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	sequence := make(chan bool)
	done := make(chan bool)

	go loopDiagonally(matrix, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
}

func readChannels(seq chan bool, done chan bool) bool {
	select {
	case <-seq:
		return true
	case <-done:
		return false
	}
}