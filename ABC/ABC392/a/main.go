package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var A [3]int
	if sc.Scan() {
		tmp := sc.Text()
		fmt.Sscanf(tmp, "%d %d %d", &A[0], &A[1], &A[2])
	}
	for i:=0; i<2; i++ {
		if A[i] > A[i+1] {
			A[i], A[i+1] = A[i+1], A[i]
		}
		if A[0] * A[1] == A[2] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}