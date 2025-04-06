package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if 400%a == 0 {
		fmt.Println(400 / a)
	} else {
		fmt.Println(-1)
	}
}
