package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scanf("%f", &x)
	if x >= 38.0 {
		fmt.Println(1)
		return
	} else if x >= 37.5 {
		fmt.Println(2)
		return
	} else {
		fmt.Println(3)
		return
	}
}
