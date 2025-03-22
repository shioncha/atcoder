package main

import (
	"fmt"
)

func main() {
	var a [7]int
	var b [13]int
	var f int

	for i := 0; i < 7; i++ {
		fmt.Scanf("%d", &a[i])
		b[a[i]-1]++
	}

	f = -1
	for i := 0; i < 13; i++ {
		if b[i] >= 3 {
			f = i
		}
	}
	if f == -1 {
		fmt.Println("No")
		return
	}
	for i := 0; i < 13; i++ {
		if b[i] >= 2 && i != f {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
