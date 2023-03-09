package luhngo

import (
	"errors"
	"strconv"
)

func GenerateDigit(input string) (string, error) {
	r, err := remainder(input)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(10 - r), nil
}

func Verify(input string) error {
	r, err := remainder(input)
	if err != nil {
		return err
	}
	if r != 0 {
		return errors.New("invalid input")
	}
	return nil
}

func rtoi(r rune) (int, error) {
	if r < '0' || r > '9' {
		return 0, errors.New("invalid r")
	}
	return int(r) - '0', nil
}

func decideNumber(i, b int) int {
	if i%2 == 1 {
		// index is odd number
		return b
	}
	// index is even number
	b *= 2
	return b%10 + b/10
}

func remainder(input string) (int, error) {
	sum := 0
	for i, s := range input {
		i++
		si, err := rtoi(s)
		if err != nil {
			return 0, err
		}
		sum += decideNumber(i, si)
	}
	return sum % 10, nil
}
