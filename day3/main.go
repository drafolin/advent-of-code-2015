package main

import (
	"fmt"
	"os"
	"slices"
)

type Coordinate struct {
	X, Y int
}

func main() {
	f, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	visited, currentPos, willMove := []Coordinate{{0, 0}}, [2]Coordinate{{0, 0}, {0, 0}}, 0

	for _, move := range f {
		switch move {
		case 'v':
			currentPos[willMove].Y--
		case '^':
			currentPos[willMove].Y++
		case '<':
			currentPos[willMove].X--
		case '>':
			currentPos[willMove].X++
		default:
			continue
		}

		if !slices.Contains(visited, currentPos[willMove]) {
			visited = append(visited, currentPos[willMove])
		}

		willMove ^= 1
	}

	fmt.Println("visited houses: ", len(visited))
}
