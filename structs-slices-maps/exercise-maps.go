package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for i := 0; i < len(words); i++ {
		result[words[i]]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
