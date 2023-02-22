package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

type Symbol struct {
	Value    rune
	IsLetter bool
	IsDigit  bool
	IsSlash  bool
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var sb strings.Builder
	var prev Symbol
	for i, val := range input {
		current := Symbol{
			Value: val,
		}
		isLast := i == len(input)-1

		if unicode.IsLetter(current.Value) {
			if prev.Value > 0 {
				sb.WriteRune(prev.Value)
			}
			if isLast {
				sb.WriteRune(current.Value)
			}

			prev = current
		}
		if unicode.IsDigit(current.Value) {
			if prev.Value > 0 {
				for i := 0; i < int(current.Value-'0'); i++ {
					sb.WriteRune(prev.Value)
				}
				prev.Value = 0
			} else {
				return "", ErrInvalidString
			}
		}
	}

	return sb.String(), nil
}
