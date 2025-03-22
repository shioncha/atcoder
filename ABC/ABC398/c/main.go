package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	var max int
	list := make([]int, 0, 3e5)
	list2 := make([]int, 0, 3e5)
	m := make(map[int]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if m[a[i]] == 0 {
			m[a[i]] = i + 1
			list = append(list, a[i])
		} else {
			m[a[i]] = -1
		}
	}

	for i := 0; i < len(list); i++ {
		if m[list[i]] == -1 {
			continue
		}
		list2 = append(list2, list[i])
	}
	if len(list2) == 0 {
		fmt.Println(-1)
		return
	}

	sort.Ints(list2)
	max = list2[len(list2)-1]
	fmt.Println(m[max])
}
