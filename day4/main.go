package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	in, salt := "iwrupvqb", 0

	for {
		salt++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", in, salt))))
		if strings.HasPrefix(hash, "000000") {
			break
		}
	}

	fmt.Println("the smallest salt value is", salt)
}
