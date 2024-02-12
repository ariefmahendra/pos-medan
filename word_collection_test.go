package pos_medan

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestWordCollection_success(t *testing.T) {
	words := []string{"Oke", "One"}
	expected := "Oekn"
	result := wordCollection(words)

	assert.Equal(t, expected, result)
}

func TestWordCollection_success_1(t *testing.T) {
	words := []string{"Abc", "bCd"}
	expected := "bACcd"
	result := wordCollection(words)

	assert.Equal(t, expected, result)
}

func TestWordCollection_success_2(t *testing.T) {
	words := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	expected := "aenrktipBDPTUdmosu"
	result := wordCollection(words)

	assert.Equal(t, expected, result)
}

func wordCollection(words []string) string {
	charCollection := make(map[rune]int)
	for _, word := range words {
		for _, char := range word {
			charCollection[char]++
		}
	}

	var charSorted []rune
	for char := range charCollection {
		charSorted = append(charSorted, char)
	}

	sort.Slice(charSorted, func(i, j int) bool {
		if charCollection[charSorted[i]] == charCollection[charSorted[j]] {
			return charSorted[i] < charSorted[j]
		}

		return charCollection[charSorted[i]] > charCollection[charSorted[j]]
	})

	var result []string
	for _, item := range charSorted {
		result = append(result, string(item))
	}

	resultStr := strings.Join(result, "")
	return resultStr
}
