package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("day8/input.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	for _, line := range strings.Split(string(f), "\n") {
		if line == "" {
			continue
		}
		fmt.Println(line)
		result := strconv.Quote(line)

		total += len(result) - len(line)
	}

	fmt.Println(total)
}
