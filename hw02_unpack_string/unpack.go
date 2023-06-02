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

func (s *Symbol) escape() {
	s.IsSlash = false
	s.IsDigit = false
	s.IsLetter = true
}

func (s *Symbol) setSlash() {
	s.Value = 0
	s.IsSlash = true
	s.IsDigit = false
	s.IsLetter = false
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	var prev Symbol
	for i, currRune := range input {
		current := Symbol{
			Value:    currRune,
			IsLetter: unicode.IsLetter(currRune),
			IsDigit:  unicode.IsDigit(currRune),
			IsSlash:  currRune == '\\',
		}
		isLast := i == len(input)-1

		if isLast && current.IsSlash && !prev.IsSlash {
			return "", ErrInvalidString
		}

		if prev.IsSlash {
			current.escape()
		}

		if current.IsLetter {
			if prev.Value > 0 {
				result.WriteRune(prev.Value)
			}
			if isLast {
				result.WriteRune(current.Value)
			}

			prev = current
		}
		if current.IsDigit {
			if prev.Value > 0 {
				for i := 0; i < int(current.Value-'0'); i++ {
					result.WriteRune(prev.Value)
				}
				prev.Value = 0
			} else {
				return "", ErrInvalidString
			}
		}

		if current.IsSlash {
			if prev.Value > 0 {
				result.WriteRune(prev.Value)
			}

			prev.setSlash()
		}
	}

	return result.String(), nil
}
