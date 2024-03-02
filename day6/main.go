package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Op int64

const (
	TURN_ON Op = iota
	TURN_OFF
	TOGGLE
)

type Coordinate struct {
	X, Y int
}

type Operation struct {
	from, to Coordinate
	op       Op
}

func main() {
	f, err := os.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}

	operations := ParseInput(string(f))

	var lights [1000][1000]int = [1000][1000]int{}

	for y := range lights {
		lights[y] = [1000]int{}
		for x := range lights[y] {
			lights[y][x] = 0
		}
	}

	for _, op := range operations {
		for y := op.from.Y; y <= op.to.Y; y++ {
			for x := op.from.X; x <= op.to.X; x++ {
				switch op.op {
				case TURN_ON:
					lights[y][x] = lights[y][x] + 1
				case TURN_OFF:
					lights[y][x] = lights[y][x] - 1
					if lights[y][x] < 0 {
						lights[y][x] = 0
					}
				case TOGGLE:
					lights[y][x] = lights[y][x] + 2
				}
			}
		}
	}

	total := 0
	for _, row := range lights {
		for _, light := range row {
			total += light
		}
	}

	fmt.Println("brightness level is of", total)
}

var exp = Assert(regexp.Compile("([a-z ]*) ([0-9]{1,3}),([0-9]{1,3}) through ([0-9]{1,3}),([0-9]{1,3})"))

func ParseInput(f string) []Operation {
	res := []Operation{}
	for _, s := range strings.Split(f, "\n") {
		matches := exp.FindStringSubmatch(s)
		if matches == nil {
			continue
		}

		fromX, fromY, toX, toY :=
			Assert(strconv.Atoi(matches[2])),
			Assert(strconv.Atoi(matches[3])),
			Assert(strconv.Atoi(matches[4])),
			Assert(strconv.Atoi(matches[5]))

		var operation Op

		switch matches[1] {
		case "turn off":
			operation = TURN_OFF
		case "turn on":
			operation = TURN_ON
		case "toggle":
			operation = TOGGLE
		}

		res = append(res, Operation{
			from: Coordinate{fromX, fromY},
			to:   Coordinate{toX, toY},
			op:   operation,
		})
	}
	return res
}

func Assert[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func Count[T comparable](s []T, ex T) int {
	cnt := 0

	for _, v := range s {
		if v == ex {
			cnt++
		}
	}
	return cnt
}
