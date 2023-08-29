package tourOfGo

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := make(map[string]int)
	for _, str := range arr {
		m[str] += 1
	}
	return m
}

func Maps() {
	wc.Test(WordCount)
}
