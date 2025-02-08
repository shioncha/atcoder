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
	var N int

	buf := make([]byte, 6*1024*1024)
	sc.Buffer(buf, 6*1024*1024)
	if sc.Scan() {
		tmp := sc.Text()
		N, _ = strconv.Atoi(tmp)
	}
	P := make([]int, N)
	Q := make([]int, N)
	if sc.Scan() {
		strs := sc.Text()
		str := strings.Split(strs, " ")
		for i:=0; i<N; i++ {
			P[i], _ = strconv.Atoi(str[i])
		}
	}
	if sc.Scan() {
		strs := sc.Text()
		str := strings.Split(strs, " ")
		for i:=0; i<N; i++ {
			Q[i], _ = strconv.Atoi(str[i])
		}
	}

	mapPQ := make(map[int]int)
	for i:=0; i<N; i++ {
		mapPQ[Q[i]-1] = P[i]
	}

	output := make([]string, N)
	for i:=0; i<N; i++ {
		output[i] = strconv.Itoa(Q[mapPQ[i]-1])
	}
	fmt.Println(strings.Join(output, " "))
}