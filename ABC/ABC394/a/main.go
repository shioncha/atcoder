package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Scan(&s)
	n := strings.Count(s, "2")

	for i := 0; i < n; i++ {
		fmt.Print("2")
	}
	fmt.Print("\n")
}
