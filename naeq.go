package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Convert the arguments to a single string
	phrase := strings.ToLower(strings.Join(os.Args[1:], ""))

	value := 0
	for _, c := range phrase {
		if unicode.IsLower(c) {
			value += int(c - 'a') * 19 % 26 + 1
		}
	}
	fmt.Println(value)
}
