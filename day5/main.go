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
		if strings.Count(s, "a")+
			strings.Count(s, "e")+
			strings.Count(s, "i")+
			strings.Count(s, "o")+
			strings.Count(s, "u") < 3 {
			continue
		}

		if !HasDuplicates(s) {
			continue
		}

		if strings.Contains(s, "ab") ||
			strings.Contains(s, "cd") ||
			strings.Contains(s, "pq") ||
			strings.Contains(s, "xy") {
			continue
		}

		count++
	}

	fmt.Println(count, "strings are nice")
}

func HasDuplicates(s string) bool {
	for i, c := range s[1:] {
		if c == rune(s[i]) {
			return true
		}
	}
	return false
}
