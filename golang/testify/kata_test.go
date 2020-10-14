package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassingRomanNumberalDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// single character roman numerals
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		// additive roman numerals
		{"II", 2},
		{"XX", 20},
		{"MDCLXVI", 1666},      // full additive roman numeral
		{"MMMCCCXXXIII", 3333}, // “pathological” additive roman numeral
		// subtractive roman numerals (base cases)
		{"IV", 4},
		{"IX", 9},
		{"XL", 40},
		{"XC", 90},
		{"CD", 400},
		{"CM", 900},
		{"MMCDLXVII", 2467},
		// pathological cases
		{"MMMCMXCIX", 3999},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q is %d", tt.input, tt.expected), func(t *testing.T) {
			got, err := RomanNumberalDecode(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestFailingRomanNumberalDecode(t *testing.T) {
	inputs := []string{
		// bad ordering
		"IM",
		"VL",
		"MMMCDVLXII",
		// too many chars
		"IIII",
		"VV",
		"XXXX",
		"LL",
		"CCCC",
		"DD",
		"MMMM",
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%q should error", input), func(t *testing.T) {
			_, err := RomanNumberalDecode(input)
			assert.Error(t, err)
		})
	}
}
