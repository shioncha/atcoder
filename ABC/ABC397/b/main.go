package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	var cnt int
	cnt = 0
	for i := 0; i < len(s); i++ {
		if ((i + cnt) % 2 == 0 && s[i] == 'i') || ((i + cnt) % 2 == 1 && s[i] == 'o') {
			continue
		}
		// fmt.Println(i+cnt%2, s[i])
		cnt++
	}
	if s[len(s)-1] == 'i' {
		cnt++
	}
	fmt.Println(cnt)
}