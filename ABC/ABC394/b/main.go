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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i] = sc.Text()
	}

	for i := 0; i < n - 1; i++ {
		for j := i; j < n; j++ {
			if len(s[i]) > len(s[j]) {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(s[i])
	}
}
