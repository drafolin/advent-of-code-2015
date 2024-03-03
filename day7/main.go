package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func exec(command string, args ...uint16) uint16 {
	switch command {
	case "NOT":
		return ^args[0]
	case "AND":
		return args[0] & args[1]
	case "OR":
		return args[0] | args[1]
	case "LSHIFT":
		return args[0] << args[1]
	case "RSHIFT":
		return args[0] >> args[1]
	case "":
		return args[0]
	}
	return 0
}

type Wire struct {
	Args      []string // may be either id or number litteral
	Operation string
	value     int32
}

var wireMap = make(map[string]Wire)

func getVal(s string) uint16 {
	v, err := strconv.Atoi(s)
	if err == nil {
		return uint16(v)
	}
	wire := wireMap[s]
	if wire.value != -1 {
		return uint16(wire.value)
	}
	wire.value = int32(calc(wire))
	wireMap[s] = wire
	return uint16(wire.value)
}

func calc(wire Wire) uint16 {
	args := []uint16{}
	for _, arg := range wire.Args {
		args = append(args, getVal(arg))
	}
	return exec(wire.Operation, args...)
}

func main() {

	f, err := os.ReadFile("day7/input.txt")

	if err != nil {
		panic(err)
	}
	s := string(f)

	exp := regexp.MustCompile(`^((?:[a-z]+|[0-9]+)?) ?([A-Z]*) ?((?:[a-z]+|[0-9]+)?) -> ([a-z]*)$`)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		fmt.Println(line)
		match := exp.FindStringSubmatch(line)
		fmt.Println(match)
		id := match[4]
		wire := Wire{Operation: match[2], value: -1}
		if match[1] != "" {
			wire.Args = append(wire.Args, match[1])
		}
		if match[3] != "" {
			wire.Args = append(wire.Args, match[3])
		}
		wireMap[id] = wire
	}
	fmt.Println(getVal("a"))
}
