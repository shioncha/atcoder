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
	var n, m int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &n, &m)
	k := make([]int, m)
	// a := make([][]int, m)
	shokuzai := make([][]int, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		tmp := strings.Split(sc.Text(), " ")
		k[i], _ = strconv.Atoi(tmp[0])
		tmpa := make([]int, k[i])
		for j := 0; j < k[i]; j++ {
			tmpa[j], _ = strconv.Atoi(tmp[j+1])
			shokuzai[tmpa[j]-1] = append(shokuzai[tmpa[j]-1], i)
		}
		// a[i] = tmpa
	}
	b := make([]int, n)
	sc.Scan()
	tmp := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		b[i], _ = strconv.Atoi(tmp[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		bi := b[i]-1
		for j := 0; j < len(shokuzai[bi]); j++ {
			k[shokuzai[bi][j]]--
			if k[shokuzai[bi][j]] == 0 {
				ans++
			}
		}
		
	}
}
