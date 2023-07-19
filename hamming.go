package exercism

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strings must be of same length")
	}

	length := len(a)

	runeA := []rune(a)
	runeB := []rune(b)

	diff := 0
	for i := 0; i < length; i++ {
		if runeA[i] != runeB[i] {
			diff++
		}
	}

	return diff, nil
}
