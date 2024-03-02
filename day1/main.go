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
	var printedBasement bool

	for i, op := range in {
		if op == '(' {
			lvl++
		} else if op == ')' {
			lvl--
		}
		if lvl < 0 && !printedBasement {
			fmt.Println("entered basement: ", i+1)
			printedBasement = true
		}
	}
	fmt.Println("ends up on lvl ", lvl)
}
