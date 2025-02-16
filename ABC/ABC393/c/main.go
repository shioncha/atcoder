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
	strs := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(strs[0])
	m, _ := strconv.Atoi(strs[1])
	u := make([]int, m)
	v := make([]int, m)
	for i:=0; i<m; i++ {
		sc.Scan()
		strs = strings.Split(sc.Text(), " ")
		u[i], _ = strconv.Atoi(strs[0])
		v[i], _ = strconv.Atoi(strs[1])
	}
	
	graph := make([]map[int]int, n)
	for i:=0; i<n; i++ {
		graph[i] = make(map[int]int)
	}
	
	cnt := 0
	for i:=0; i<m; i++ {
		ui := u[i] - 1
		vi := v[i] - 1
		if ui < vi {
			ui, vi = vi, ui
		}
		if u[i] == v[i] || graph[ui][vi] == 1 {
			cnt++
		}
		graph[ui][vi] = 1
	}
	fmt.Println(cnt)
}
