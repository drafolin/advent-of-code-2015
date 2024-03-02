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

	visited, currentPos := []Coordinate{{0, 0}}, Coordinate{0, 0}

	for _, move := range f {
		switch move {
		case 'v':
			currentPos.Y--
		case '^':
			currentPos.Y++
		case '<':
			currentPos.X--
		case '>':
			currentPos.X++
		default:
			continue
		}

		if !slices.Contains(visited, currentPos) {
			visited = append(visited, currentPos)
		}
	}

	fmt.Println("visited houses: ", len(visited))
}
