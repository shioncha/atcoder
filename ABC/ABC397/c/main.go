package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	l := make(map[int]int)
	var lk, rk int
	r := make(map[int]int)
	for i := 0; i < n; i++ {
		if r[a[i]] < 1 {
			rk++
		}
		r[a[i]]++
	}

	for i := 0; i < n-1; i++ {
		r[a[i]]--
		l[a[i]]++
		if l[a[i]] < 2 {
			lk++
		}
		if r[a[i]] < 1 {
			rk--
		}
		// fmt.Println(lk, rk)
		ans = max(ans, lk+rk)
	}
	fmt.Println(ans)
}
