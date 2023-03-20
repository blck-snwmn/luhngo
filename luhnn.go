package luhngo

import (
	"errors"
	"strings"
)

func NewLuhnn(sl string) (luhnn, error) {
	if len(sl)%2 != 0 {
		return luhnn{}, errors.New("invalid input string")
	}
	return luhnn{sl: sl, n: len(sl)}, nil
}

type luhnn struct {
	sl string
	n  int
}

func (l luhnn) Generate(input string) (string, error) {
	r, err := l.remainder(input)
	if err != nil {
		return "", err
	}
	return l.ptoc(r), nil
}

func (l luhnn) Verify(input string) error {
	r, err := l.remainder(input)
	if err != nil {
		return err
	}
	if r != 0 {
		return errors.New("invalid input")
	}
	return nil
}

func (l luhnn) decideNumber(i, b int) int {
	if i%2 == 1 {
		// index is odd number
		return b
	}
	// index is even number
	b *= 2
	return b%l.n + b/l.n
}

func (l luhnn) remainder(input string) (int, error) {
	sum := 0
	for i, s := range input {
		i++ // index start 1
		si := l.ctop(s)
		sum += l.decideNumber(i, si)
	}
	return sum % l.n, nil
}

func (l luhnn) ctop(r rune) int {
	return strings.IndexRune(l.sl, r)
}

func (l luhnn) ptoc(p int) string {
	return string(l.sl[l.n-p-1])
}
