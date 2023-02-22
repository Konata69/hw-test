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
			Value:    val,
			IsLetter: unicode.IsLetter(val),
			IsDigit:  unicode.IsDigit(val),
			IsSlash:  val == []rune(`\`)[0],
		}
		isLast := i == len(input)-1

		if current.IsLetter {
			if prev.Value > 0 {
				sb.WriteRune(prev.Value)
			}
			if isLast {
				sb.WriteRune(current.Value)
			}

			prev = current
		}
		if current.IsDigit {
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
