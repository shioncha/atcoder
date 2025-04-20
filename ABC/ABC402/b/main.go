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
	var q int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &q)
	q1 := make([]int, q)
	q2 := make([]int, 0)
	for i := 0; i < q; i++ {
		sc.Scan()
		tmp := strings.Split(sc.Text(), " ")
		tmpa, _ := strconv.Atoi(tmp[0])
		q1[i] = tmpa
		if tmpa == 1 {
			tmpb, _ := strconv.Atoi(tmp[1])
			q2 = append(q2, tmpb)
		}
	}
	p := 0
	for i := 0; i < q; i++ {
		if q1[i] == 2 {
			fmt.Printf("%d\n", q2[p])
			p++
		}
	}
}
