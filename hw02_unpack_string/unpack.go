package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(a string) (string, error) {
	var str strings.Builder
	var preVal rune
	b := []rune(a)
	if len(b) == 0 {
		return "", nil
	}
	for _, val := range b {
		if preVal == 0 {
			if unicode.IsDigit(val) { // check if 1st sign is digit
				return "", ErrInvalidString
			}
			preVal = val
			continue
		}
		if unicode.IsDigit(val) && unicode.IsDigit(preVal) { // check if 2 digits side by side
			return "", ErrInvalidString
		}
		if !unicode.IsDigit(val) || unicode.IsDigit(preVal) {
			if !unicode.IsDigit(preVal) && !unicode.IsDigit(val) {
				str.WriteString(string(preVal))
			}
			preVal = val
			continue
		}
		strRepStep, _ := strconv.Atoi(string(val))
		str.WriteString(strings.Repeat(string(preVal), strRepStep))
		preVal = val
	}
	if !unicode.IsDigit(preVal) {
		str.WriteString(string(preVal))
	}
	return str.String(), nil
}
