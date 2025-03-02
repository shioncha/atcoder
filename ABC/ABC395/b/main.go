package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1e6)
	sc.Buffer(buf, len(buf))

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	ans := make([][]string, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > n+1-i {
				continue
			}
			if i % 2 == 0{
				for k := i; k < n-i; k++ {
					for l := i; l < n-i; l++ {
						ans[k][l] = "#"
					}
				}
			}
			if i % 2 == 1{
				for k := i; k < n-i; k++ {
					for l := i; l < n-i; l++ {
						ans[k][l] = "."
					}
				}
			}
		}
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(ans[i][j])
		}
		fmt.Println()
	}
}
