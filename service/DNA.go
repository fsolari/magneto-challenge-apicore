package service

import (
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/dao"
	"log"
	"time"
)

func stringToRunesArray(dna string) []rune {
	var runes []rune

	for _, rune := range dna {
		runes = append(runes, rune)
	}

	return runes
}

func generateDNAMatrix(dna []string) [][]rune {
	matrix := [][]rune{}

	for _, s := range dna {
		runes := stringToRunesArray(s)
		matrix = append(matrix, runes)
	}
	return matrix
}

func loopHorizontally(dnaMatrix [][]rune, rows int, columns int, sequence chan bool, done chan bool) {

	var x, y, i int

	for x = 0; x < columns; x++ {

		for y = 0; y < rows - 1; y++ {

			if dnaMatrix[x][y] == dnaMatrix[x][y + 1] && i < 4 {
				i++
			} else {
				i = 0
				break
			}
			if i == 3 {
				sequence <- true
			}
		}
	}
	done <- true
}

func loopVertically(dnaMatrix [][]rune, rows int, columns int, sequence chan bool, done chan bool) {
	var x, y, i int

	for y = 0; y < rows; y++ {

		for x = 0; x < columns - 1; x++ {

			if dnaMatrix[x][y] == dnaMatrix[x + 1][y] && i < 4 {
				i++
			} else {
				i = 0
				break
			}
			if i == 3 {
				sequence <- true
			}
		}
	}
	done <- true
}

func loopDiagonally(dnaMatrix [][]rune, rows int, sequence chan bool, done chan bool) {

	var x, y, i int

	for x = 0; x < rows - 1; x++ {
		for y = 0; y < rows - 1; y++ {
			if x == y {

				if dnaMatrix[x][y] == dnaMatrix[x + 1][y + 1] && i < 4 {
					i++
				} else {
					i = 0
					break
				}
				if i == 3 {
					sequence <- true
				}
			}
		}
	}
	time.Sleep(3 * time.Millisecond)
	done <- true
}

func DNATest(dna domain.DNA) (bool, error) {

	dna.Mutant = handleMatrix(dna)

	err := dao.InsertDNA(dna)
	if err != nil {
		log.Printf("Error saving DNA registry on DB %v\n", err)
	}

	return dna.Mutant, err
}

func handleMatrix(dna domain.DNA) bool {
	sequence := make(chan bool)
	done := make(chan bool)
	rows := len(dna.DNA)
	columns := len(dna.DNA[0])

	dnaMatrix := generateDNAMatrix(dna.DNA)

	go loopHorizontally(dnaMatrix, rows, columns, sequence, done)
	go loopVertically(dnaMatrix, rows, columns, sequence, done)
	go loopDiagonally(dnaMatrix, rows, sequence, done)

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
