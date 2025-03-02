package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1e6)
	sc.Buffer(buf, len(buf))

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	a := make([]int, n)
	for i, v := range s {
		a[i], _ = strconv.Atoi(v)
	}

	for i := 0; i < n - 1; i++ {
		if a[i] >= a[i+1] {
			fmt.Println("No")
			return
		}
	}
	
	fmt.Println("Yes")
}
