package main

import (
	"fmt"
	"strings"
)

func main() {
  var N, M int
  fmt.Scanf("%d %d", &N, &M)

	S := make([][]string, 0, N)
	for i:=0; i<N; i++ {
		var strs string
		fmt.Scanf("%s", &strs)
		str := strings.Split(strs, "")
		S = append(S, str)
	}
	T := make([][]string, 0, M)
	for i:=0; i<M; i++ {
		var strs string
		fmt.Scanf("%s", &strs)
		str := strings.Split(strs, "")
		T = append(T, str)
	}

	for i:=0; i<N-M+1; i++ {
		for j:=0; j<N-M+1; j++ {
			flag := false
			for k:=0; k<M; k++ {
				for l:=0; l<M; l++ {
					if S[i+k][j+l] != T[k][l] {
						flag = true
						break
					}
					if k == M-1 && l == M-1 {
						fmt.Println(i+1, j+1)
						return
					}
				}
				if flag {
					break
				}
			}
		}
	}
}
