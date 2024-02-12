package pos_medan

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSumChar_success(t *testing.T) {
	word := "We Always Mekar"
	expected := map[string]int{
		"w": 2,
		"e": 2,
		"a": 3,
		"l": 1,
		"y": 1,
		"s": 1,
		"m": 1,
		"k": 1,
		"r": 1,
	}

	result := sumChar(word)

	assert.Equal(t, expected, result)
}

func TestSumChar_success_1(t *testing.T) {
	word := "coding is fun"
	expected := map[string]int{
		"c": 1,
		"o": 1,
		"d": 1,
		"i": 2,
		"n": 2,
		"g": 1,
		"s": 1,
		"f": 1,
		"u": 1,
	}

	result := sumChar(word)

	assert.Equal(t, expected, result)
}

func sumChar(str string) map[string]int {
	strLowerCase := strings.ToLower(str)

	strArr := strings.Split(strLowerCase, "")

	sumCharUnique := make(map[string]int)
	for _, item := range strArr {
		if item == " " {
			continue
		} else if sumCharUnique[item] != 0 {
			sumCharUnique[item]++
		} else {
			sumCharUnique[item] = 1
		}
	}

	return sumCharUnique
}
