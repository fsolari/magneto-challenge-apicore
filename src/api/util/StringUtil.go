package util

import "strings"

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