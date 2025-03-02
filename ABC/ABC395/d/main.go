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
	s1 := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s1[0])
	m, _ := strconv.Atoi(s1[1])
	op1 := make([]int, m)
	op2 := make([]int, m)
	op3 := make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		op1[i], _ = strconv.Atoi(s[0])
		op2[i], _ = strconv.Atoi(s[1])
		if op1[i] < 3 {
			op3[i], _ = strconv.Atoi(s[2])
		}
	}

	s := make([]int, n)
	h := make([]int, n)
	mm := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		s[i] = 1
		h[i] = i
		mm[i] = []int{i}
	}
	for i := 0; i < m; i++ {
		switch op1[i] {
		case 1:
			s[op2[i]-1]--
			s[op3[i]-1]++
			for j := 0; j < len(mm[h[op2[i]-1]]); j++ {
				if mm[h[op2[i]-1]][j] == op2[i]-1 {
					mm[h[op2[i]-1]] = append(mm[h[op2[i]-1]][:j], mm[h[op2[i]-1]][j+1:]...)
					break
				}
			}
			h[op2[i]-1] = op3[i]-1
			mm[op3[i]-1] = append(mm[op3[i]-1], op2[i]-1)
		case 2:
			mm[op2[i]-1], mm[op3[i]-1] = mm[op3[i]-1], mm[op2[i]-1]
			for _, v := range mm[op2[i]-1] {
				h[v] = op2[i]-1
			}
			for _, v := range mm[op3[i]-1] {
				h[v] = op3[i]-1
			}
			s[op2[i]-1], s[op3[i]-1] = s[op3[i]-1], s[op2[i]-1]
		case 3:
			fmt.Println(h[op2[i]-1]+1)
		}
	}
}
