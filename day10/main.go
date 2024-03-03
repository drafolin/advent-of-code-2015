package main

import "fmt"

func main() {
	const runs = 40
	input := "3113322113"

	for i := 0; i < runs; i++ {
		res, consecutive, current := "", 0, rune(input[0])
		for _, char := range input {
			if current != char {
				res += fmt.Sprint(consecutive) + string(current)
				current, consecutive = char, 1
			} else {
				consecutive++
			}
		}
		res += fmt.Sprint(consecutive) + string(current)
		input = res
	}

	fmt.Println(len(input))
}
