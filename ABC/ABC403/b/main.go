package main

import (
	"fmt"
)

func main() {
	var t, u string
	fmt.Scanf("%s\n", &t)
	fmt.Scanf("%s", &u)

	flag := true
	for i := 0; i < len(t)-len(u)+1; i++ {
		flag = true
		for j := 0; j < len(u); j++ {
			if t[i+j] == '?' {
				continue
			}
			if t[i+j] != u[j] {
				flag = false
				break
			}
		}

		if flag {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
