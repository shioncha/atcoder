package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n%2 == 0{
		for i := 0; i < n; i++ {
			if i == n / 2 - 1 || i == n / 2 {
				fmt.Print("=")
				continue
			}
			fmt.Print("-")
		}
	} else {
		for i := 0; i < n; i++ {
			if i == n / 2 {
				fmt.Print("=")
				continue
			}
			fmt.Print("-")
		}
	}
	fmt.Println()
}
