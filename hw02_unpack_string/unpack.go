package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var sb strings.Builder
	var prev rune = 0
	for i, val := range input {
		isLast := i == len(input)-1

		if unicode.IsLetter(val) {
			if prev > 0 {
				sb.WriteRune(prev)
			}
			if isLast {
				sb.WriteRune(val)
			}

			prev = val
		}
		if unicode.IsDigit(val) {
			if prev > 0 {
				for i := 0; i < int(val-'0'); i++ {
					sb.WriteRune(prev)
				}
				prev = 0
			} else {
				return "", ErrInvalidString
			}
		}
	}

	return sb.String(), nil
}
