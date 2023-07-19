package exercism

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char > 'Z' || char < 'A' {
		return "", errors.New("Char out of range")
	}

	length := int(2*(char-'A') + 1)
	matrix := make([][]byte, length)

	base := make([]byte, length)
	current := make([]byte, length)

	for i := 0; i < length; i++ {
		base[i] = byte(' ')
	}

	for i := 0; i <= length/2; i++ {
		copy(current, base)
		current[length/2-i] = 'A' + byte(i)
		current[length/2+i] = 'A' + byte(i)

		matrix[i] = make([]byte, length)
		copy(matrix[i], current)

		if i < length/2 {
			matrix[length-i-1] = make([]byte, length)
			copy(matrix[length-i-1], current)
		}
	}

	var builder strings.Builder

	for idx, str := range matrix {
		builder.WriteString(string(str))
		if idx < length-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String(), nil
}
