package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	total := 0
	for _, l := range input {
		dim := [3]string(strings.Split(l, "x"))
		lS, wS, hS := dim[0], dim[1], dim[2]

		l, lerr := strconv.Atoi(lS)
		w, werr := strconv.Atoi(wS)
		h, herr := strconv.Atoi(hS)

		for _, err := range []error{lerr, werr, herr} {
			if err != nil {
				panic(err)
			}
		}

		sides := [3]int{l * w, l * h, w * h}
		slack := slices.Min(sides[:])
		total += sides[0]*2 +
			sides[1]*2 +
			sides[2]*2 +
			slack
	}
	fmt.Println(total)
}
