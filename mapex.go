package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ws := strings.Fields(s)
	wm := make(map[string]int)
	for _, v := range ws {
		_, ok := wm[v]
		if ok {
			wm[v] = wm[v] + 1
		} else {
			wm[v] = 1
		}

	}
	return wm
}

func main() {
	wc.Test(WordCount)
}
