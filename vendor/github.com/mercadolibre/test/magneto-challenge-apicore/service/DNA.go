package service

import (
	"github.com/mercadolibre/test/magneto-challenge-apicore/domain"
	"fmt"
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


func loopHorizontally(dnaMatrix [][]rune, rows int, columns int) bool{

	var x, y, i int

	for x = 0; x < columns; x++ {

		for y = 0; y < rows - 1; y++ {

			if dnaMatrix[x][y] == dnaMatrix[x][y + 1] && i < 4 {
				i++
			} else {
				i = 0
				break
			}
			if i == 3{
				return true
			}
		}
	}
	return false
}

func loopVertically(dnaMatrix [][]rune, rows int, columns int) bool {
	var x, y, i int

	for y = 0; y < rows; y++ {

		for x = 0; x < columns - 1; x++ {

			if dnaMatrix[x][y] == dnaMatrix[x + 1][y] && i < 4 {
				i++
			} else {
				i = 0
				break
			}
			if i == 3{
				return true
			}
		}
	}
	return false
}

func loopDiagonally(dnaMatrix [][]rune, rows int, columns int) bool{

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
				if i == 3{
					return true
				}
			}
		}
	}
	return false
}


func DNATest(dna domain.DNA) (bool, error) {

	rows := len(dna.DNA)
	columns := len(dna.DNA[0])

	dnaMatrix := generateDNAMatrix(dna.DNA)

	fmt.Print(dnaMatrix)


	if loopHorizontally(dnaMatrix, rows, columns) || loopVertically(dnaMatrix, rows, columns) || loopDiagonally(dnaMatrix, rows, columns){

		return true, nil
	}

	return false, nil
}