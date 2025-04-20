package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	pattern := `([A-Z]?)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(s, -1)
	for _, match := range matches {
		if match != "" {
			fmt.Print(match)
		}
	}
}
