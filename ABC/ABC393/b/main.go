package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var cnt int = 0

	fmt.Scan(&s)

	if !strings.Contains(s, "A") || !strings.Contains(s, "B") || !strings.Contains(s, "C") {
		fmt.Println(0)
		return
	}

	for i:=0; i<len(s); i++ {
		if s[i] != 'A' {
			continue
		}
		for j:=i+1; j<len(s); j++ {
			if s[j] != 'B' {
				continue
			}
			for k:=j+1; k<len(s); k++ {
				if s[k] != 'C' {
					continue
				}
				if j - i == k - j {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}