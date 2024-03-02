package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("day5/input.txt")
	if err != nil {
		panic(err)
	}

	s, count := strings.Split(string(f), "\n"), 0

	for _, s := range s {
		if len(s) < 2 {
			continue
		}

		if !HasDoubleDupes(s) {
			continue
		}

		if !HasDuplicates(s) {
			continue
		}

		count++
	}

	fmt.Println(count, "strings are nice")
}

func HasDoubleDupes(s string) bool {
	for i, c := range s[1:] {
		if strings.Count(s, string([]rune{rune(s[i]), c})) >= 2 {
			return true
		}
	}
	return false
}

func HasDuplicates(s string) bool {
	for i, c := range s[2:] {
		if c == rune(s[i]) {
			return true
		}
	}
	return false
}
