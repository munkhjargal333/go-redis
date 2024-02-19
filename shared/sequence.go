package shared

import (
	"errors"
	"unicode"
)

func next(ch rune) (string, error) {
	if !unicode.IsDigit(ch) && !unicode.IsLetter(ch) {
		return "_", nil
	}

	if ch == 'Z' {
		return "a", nil
	} else if ch == 'z' {
		return "1", nil
	} else if ch == '9' {
		return "", errors.New("is ended")
	} else {
		return string(ch + 1), nil
	}
}

func GenerateSeq(currentSeq string) string {
	if currentSeq == "" {
		return "AA"
	}

	if char, err := next(rune(currentSeq[1])); err != nil {
		if char, err := next(rune(currentSeq[0])); err != nil {
			return err.Error()
		} else {
			return char + "A"
		}
	} else {
		return string(currentSeq[0]) + char
	}
}
