package service

import (
	"strings"
	"errors"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertAuto(input string) (string, error) {

	input = strings.TrimSpace(input) 

	if input == "" {
		return "", errors.New("пустая строка")
	}

	morseChars := ".- "

	for _, r := range input {
		if !strings.ContainsRune(morseChars, r) {
			return morse.ToMorse(input), nil
		}
	}

	return morse.ToText(input), nil
 
}

