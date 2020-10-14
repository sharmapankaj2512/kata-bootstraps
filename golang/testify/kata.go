package kata

import (
	"errors"
)

func romanNumeral(r rune) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}

	return 0
}

var canPrecede = map[rune]rune{
	'V': 'I',
	'X': 'I',
	'L': 'X',
	'C': 'X',
	'D': 'C',
	'M': 'C',
}

func RomanNumberalDecode(s string) (int, error) {
	var sum int

	runes := []rune(s)

	for i, r := range runes {
		here := romanNumeral(r)
		next := r

		if i < len(runes)-1 {
			next = runes[i+1]
		}

		if here < romanNumeral(next) {
			if canPrecede[next] != r {
				return 0, errors.New("bad input")
			}

			sum -= here
		} else {
			sum += here
		}
	}

	return sum, nil
}
