package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var m, iii, jjj int
	for i := 0; i < n; i++ {
		iii = i * i * i
		if iii < n {
			continue
		}
		for j := 0; j < i; j++ {
			jjj = j * j * j
			m = iii - jjj
			if m == n {
				fmt.Println(i, j)
				return
			}
		}
	}

	fmt.Println(-1)
}