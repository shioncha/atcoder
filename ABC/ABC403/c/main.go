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
	sc.Buffer(make([]byte, 1e8), 1e8)
	var n, m, q int
	sc.Scan()
	tmp := strings.Split(sc.Text(), " ")
	n, _ = strconv.Atoi(tmp[0])
	m, _ = strconv.Atoi(tmp[1])
	q, _ = strconv.Atoi(tmp[2])
	_ = m

	graph := make([]map[int]int, n)
	for i:=0; i<n; i++ {
		graph[i] = make(map[int]int)
	}
	all := make([]int, n)

	output := make([]string, q)

	var q1, q2, q3 int
	for i := 0; i < q; i++ {
		sc.Scan()
		tmp := strings.Split(sc.Text(), " ")
		q1, _ = strconv.Atoi(tmp[0])
		q2, _ = strconv.Atoi(tmp[1])
		q2 -= 1
		switch q1 {
		case 1:
			q3, _ = strconv.Atoi(tmp[2])
			q3 -= 1
			graph[q2][q3] = 1
		case 2:
			all[q2] = 1
		case 3:
			q3, _ = strconv.Atoi(tmp[2])
			q3 -= 1
			if all[q2] == 1 {
				output[i] = "Yes"
				continue
			}
			if graph[q2][q3] == 1 {
				output[i] = "Yes"
				continue
			}
			output[i] = "No"
		}
	}

	for i := 0; i < q; i++ {
		if output[i] != "" {
			fmt.Println(output[i])
		}
	}
}
