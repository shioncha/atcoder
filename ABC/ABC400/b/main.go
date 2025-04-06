package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	v := 1
	ans := 1
	for i := 1; i <= m; i++ {
		v *= n
		ans += v
		if ans > 1000000000 {
			fmt.Println("inf")
			return
		}
	}
	fmt.Println(ans)
}
