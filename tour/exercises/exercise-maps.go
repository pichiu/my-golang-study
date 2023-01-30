package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
	for _, v := range(strings.Fields(s)) {
	    count := m[v]
		m[v] = count + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
