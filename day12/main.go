package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day12/input.json")

	if err != nil {
		panic(err)
	}

	var data interface{}

	json.Unmarshal(file, &data)

	fmt.Println(recurseTotal(data))
}

func recurseTotal(v interface{}) int {
	total := 0
	switch i := v.(type) {
	case map[string]interface{}:
		for _, v := range i {
			total += recurseTotal(v)
		}
	case []interface{}:
		for _, v := range i {
			total += recurseTotal(v)
		}
	case float64:
		total += int(i)
	}
	return total
}
