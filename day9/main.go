package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var distances = map[string]map[string]int{}

var regex = regexp.MustCompile(`([A-Za-z]*) to ([A-Za-z]*) = (\d*)`)

func getMin(used []string, last string) int {
	if len(used) == len(distances) {
		return 0
	}

	minVal := math.MaxInt
	for key := range distances {
		if slices.Contains(used, key) {
			continue
		}

		val := 0
		if last != "" {
			val += distances[last][key]
		}

		minVal = min(minVal, val+getMin(append(used, key), key))
	}
	return minVal
}

func getMax(used []string, last string) int {
	if len(used) == len(distances) {
		return 0
	}

	maxVal := math.MinInt
	for key := range distances {
		if slices.Contains(used, key) {
			continue
		}

		val := 0
		if last != "" {
			val += distances[last][key]
		}

		maxVal = max(maxVal, val+getMax(append(used, key), key))
	}
	return maxVal
}

func main() {
	f, err := os.ReadFile("day9/input.txt")
	if err != nil {
		panic(err)
	}

	fs := string(f)
	lines := strings.Split(fs, "\n")
	for _, line := range lines[:len(lines)-1] {
		match := regex.FindStringSubmatch(line)
		n1, n2, dS := match[1], match[2], match[3]
		d, err := strconv.Atoi(dS)
		if err != nil {
			panic(err)
		}

		if distances[n1] == nil {
			distances[n1] = map[string]int{}
		}
		distances[n1][n2] = d

		if distances[n2] == nil {
			distances[n2] = map[string]int{}
		}
		distances[n2][n1] = d
	}

	fmt.Println(getMin([]string{}, ""))
	fmt.Println(getMax([]string{}, ""))
}
