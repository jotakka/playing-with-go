package exercism

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var sb strings.Builder

	switch {
	case number%3 == 0, number%5 == 0, number%7 == 0:
		if number%3 == 0 {
			sb.WriteString("Pling")
		}
		if number%5 == 0 {
			sb.WriteString("Plang")
		}
		if number%7 == 0 {
			sb.WriteString("Plong")
		}
	default:
		return strconv.Itoa(number)
	}
	return sb.String()
}
