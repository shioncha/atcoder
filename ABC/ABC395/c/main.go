package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// buf := make([]byte, 1e10)
	// sc.Buffer(buf, len(buf))

	// sc.Scan()
	// n, _ := strconv.Atoi(sc.Text())
	// sc.Scan()
	// s := strings.Split(sc.Text(), " ")
	var n int
	// var s string
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i, _ := range a {
		// a[i], _ = strconv.Atoi(v)
		fmt.Scanf("%d", &a[i])
	}

	m := make(map[int]int)
	mm := make(map[int]int)
	max := 1000000
	min := max
	if n == 1 {
		fmt.Println(-1)
		return
	}
	for i := 0; i < n; i++ {
		m[a[i]]++
		if m[a[i]] > 1 {
			sa := i - mm[a[i]] + 1
			if sa < min {
				min = sa
			}
		}
		mm[a[i]] = i
	}
	if min == max {
		fmt.Println(-1)
		return
	}
	fmt.Println(min)
}
