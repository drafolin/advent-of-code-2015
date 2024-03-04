package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dlclark/regexp2"
)

func main() {
	password := "hxbxwxba"
	for !criterions(password) {
		password = increment(password)
	}
	fmt.Println(password)
	password = increment(password)
	for !criterions(password) {
		password = increment(password)
	}
	fmt.Println(password)
}

func increment(s string) string {
	password := []rune(s)
	slices.Reverse(password)

	lastZ := -1

	for password[lastZ+1] == 'z' {
		lastZ++
	}

	if lastZ == len(password)-1 {
		return strings.Repeat("a", len(password)+1)
	}

	for i := 0; i <= lastZ; i++ {
		password[i] = 'a'
	}

	password[lastZ+1] += 1

	slices.Reverse(password)
	return string(password)
}

func criterion1(s string) bool {
	var nextChar rune
	consecutive := 0
	for _, char := range s {
		if char == nextChar {
			consecutive++
		} else {
			consecutive = 1
		}
		if consecutive == 3 {
			return true
		}

		nextChar = char + 1
	}

	return false
}

func criterion2(s string) bool {
	return !(strings.Contains(s, "i") ||
		strings.Contains(s, "o") ||
		strings.Contains(s, "l"))
}

var c3Regex = regexp2.MustCompile(`.*(.)\1.*(.)\2.*`, 0)

func criterion3(s string) bool {
	res, _ := c3Regex.MatchString(s)
	return res
}
func criterions(s string) bool {
	return criterion1(s) && criterion2(s) && criterion3(s)
}
