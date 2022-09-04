package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range(words) {
		if _, ok := m[v]; ok == false {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
