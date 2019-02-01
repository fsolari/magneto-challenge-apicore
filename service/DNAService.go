package service

import (
	"github.com/mercadolibre/magneto-challenge-apicore/dao"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/util"
	"time"
)

func DNATest(dna domain.DNA) (bool, error) {

	err := dao.GetDNA(dna)
	if err != nil {
		return false, err
	}

	cols := len(dna.DNA[0])
	DNA := util.GenerateMatrixFromStringArray(dna.DNA)

	dna.Mutant = isMutant(DNA, cols)

	return dna.Mutant, dao.InsertDNA(dna)

}

func isMutant(DNA [][]rune, cols int) bool {

	sequence := make(chan bool)
	done := make(chan bool)

	go loopHorizontally(DNA, cols, sequence, done)
	go loopVertically(DNA, cols, sequence, done)
	go loopDiagonally(DNA, sequence, done)

	return channelReader(sequence, done)
}

func channelReader(sequence chan bool, done chan bool) bool {

	for i := 0; i <= 2; i++ {
		select {
		case <-sequence:
			if i == 1 {
				return true
			}
		case <-done:
			if i == 2 {
				return false
			}
		}
	}

	return false
}

func loopHorizontally(DNA [][]rune, cols int, sequence chan bool, done chan bool) {

	var x, y, i int

	for x = 0; x < cols; x++ {

		for y = 0; y < len(DNA) - 1; y++ {

			if checkCondition(DNA[x][y], DNA[x][y + 1], i) {
				i++
			} else {
				i = 0
				break
			}
			onSequenceSendTrue(i, 3, sequence)
		}
	}
	done <- true
}

func loopVertically(DNA [][]rune, cols int, sequence chan bool, done chan bool) {
	var x, y, i int

	for y = 0; y < len(DNA); y++ {

		for x = 0; x < cols - 1; x++ {

			if checkCondition(DNA[x][y], DNA[x + 1][y], i) {
				i++
			} else {
				i = 0
				break
			}
			onSequenceSendTrue(i, 3, sequence)
		}
	}
	time.Sleep(1 * time.Millisecond)
	done <- true
}

func loopDiagonally(DNA [][]rune, sequence chan bool, done chan bool) {

	var x, y, i int

	for x = 0; x < len(DNA) - 1; x++ {
		for y = 0; y < len(DNA) - 1; y++ {
			if x >= y {
				if checkCondition(DNA[x][y], DNA[x + 1][y + 1], i) {
					i++
				} else {
					i = 0
					break
				}
				onSequenceSendTrue(i, 3, sequence)
			}
		}
	}
	time.Sleep(1 * time.Millisecond)
	done <- true
}

func checkCondition(a rune, b rune, i int) bool {
	if a == b && i < 4 {
		return true
	}
	return false
}

func onSequenceSendTrue(i int, j int, channel chan bool) {
	if i == j {
		channel <- true
	}
}
