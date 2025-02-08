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
	var M int
	if sc.Scan() {
		tmp := sc.Text()
		fmt.Sscanf(tmp, "%d %d", &N, &M)
	}
	A := make([]int, M)
	if sc.Scan() {
		strs := sc.Text()
		str := strings.Split(strs, " ")
		for i:=0; i<M; i++ {
			A[i], _ = strconv.Atoi(str[i])
		}
	}
	var output []string
	for i:=1; i<=N; i++ {
		var flag bool = false
		for j:=0; j<M; j++ {
			if A[j] == i {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		output = append(output, strconv.Itoa(i))
	}
	fmt.Println(len(output))
	fmt.Println(strings.Join(output, " "))
}