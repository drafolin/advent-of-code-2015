package utils

import (
	"os"
	"strings"
)

func RemoveFromSlice[T comparable](s *[]T, r T) {
	res := []T{}
	for _, v := range *s {
		if v != r {
			res = append(res, v)
		}
	}
	*s = res
}

func ReadFile(path string) string {
	f, err := os.ReadFile("day10/input.txt")

	if err != nil {
		panic(err)
	}

	RemoveFromSlice(&f, '\r')
	return string(f)
}

func ReadLines(path string) []string {
	f := ReadFile(path)
	return strings.Split(f, "\n")
}
