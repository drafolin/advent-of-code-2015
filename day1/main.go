package main

import (
	"fmt"
	"os"
)

func main() {
	in, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	lvl := 0
	for _, op := range in {
		if op == '(' {
			lvl++
		} else if op == ')' {
			lvl--
		}
	}
	fmt.Println(lvl)
}
