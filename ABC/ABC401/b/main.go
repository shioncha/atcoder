package main

import (
	"fmt"
)

func main() {
  var n int
  fmt.Scanf("%d\n", &n)

  cnt := 0
	isLogin := false
	var s string
	for i:=0; i<n; i++ {
		fmt.Scanf("%s\n", &s)

		if s == "login" {
			isLogin = true
		} else if s == "logout" {
			isLogin = false
		} else if s == "private" {
			if !isLogin {
				cnt++
			}
		} else {
			continue
		}
	}

	fmt.Println(cnt)
}
