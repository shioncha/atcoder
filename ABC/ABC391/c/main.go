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
  var N, Q int
	sc.Scan()
	tmp := strings.Split(sc.Text(), " ")
	N, _ = strconv.Atoi(tmp[0])
	Q, _ = strconv.Atoi(tmp[1])
	P := make([]int, 0, N)
	H := make([]int, 0, N)
	for i:=0; i<N; i++ {
		P = append(P, i)
		H = append(H, 1)
	}
	cnt := 0
	var L, R int
	for i:=0; i<Q; i++ {
		sc.Scan()
		tmpx := strings.Split(sc.Text(), " ")
		if tmpx[0] == "2" {
			fmt.Println(cnt)
			continue
		} else if tmpx[0] == "1" {
			L, _ = strconv.Atoi(tmpx[1])
			R, _ = strconv.Atoi(tmpx[2])
			if H[P[L-1]] == 2 {
				cnt--
			}
			if H[R-1] == 1 {
				cnt++
			}
			H[P[L-1]]--;
			H[R-1]++;
			P[L-1] = R-1
		}
	}
}
