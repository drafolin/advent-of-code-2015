package main

import (
	"fmt"

	regexp "github.com/dlclark/regexp2"
)

func main() {
	const runs = 50
	input := "3113322113"
	regex := regexp.MustCompile(`(\d)\1*`, 0)

	for i := 0; i < runs; i++ {
		temp := ""
		m, _ := regex.FindStringMatch(input)
		for m != nil {
			temp += fmt.Sprintf("%d%s", len(m.String()), m.Groups()[1].String())
			m, _ = regex.FindNextMatch(m)
		}
		input = temp
		go fmt.Println(i)
	}

	fmt.Println(len(input))
}
