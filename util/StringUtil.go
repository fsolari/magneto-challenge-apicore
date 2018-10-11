package util

import (
	"strings"
	"time"
	"fmt"
	"math/rand"
	"golang.org/x/net/html/charset"
	"github.com/heroku/go-getting-started/domain"
)

func StringArrayContains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}

func stringToRuneArray(s string) []rune {
	var runes []rune

	for _, rune := range s {
		runes = append(runes, rune)
	}

	return runes
}

func GenerateMatrixFromStringArray(a []string) [][]rune {
	matrix := [][]rune{}

	for _, s := range a {
		runes := stringToRuneArray(s)
		matrix = append(matrix, runes)
	}
	return matrix
}

func JoinStringArray(arr []string) string {
	return strings.Join(arr, " ")
}


func generateRandomStringArray() []string{

	var arr []string

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 6)

	for i := range b {
		b[i] = domain.DNACharSet[seededRand.Intn(len(domain.DNACharSet))]
	}

	for i:=0; i <= 6; i++ {
		arr = append(arr, string(b))

	}

	return arr
}