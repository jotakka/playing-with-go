package exercirsm

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	orderedSubj := orderStr(subject)

	anagrams := make([]string, len(candidates))
	count := 0

	for _, str := range candidates {
		ordered := orderStr(str)
		if ordered == orderedSubj && strings.ToLower(subject) != strings.ToLower(str) {
			anagrams[count] = str
			count++
		}
	}

	return anagrams[:count]
}

func orderStr(input string) string {
	split := strings.Split(strings.ToLower(input), "")
	sort.Strings(split)
	return strings.Join(split, "")
}
