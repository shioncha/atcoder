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
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	a := strings.Split(sc.Text(), " ")

	num := "-1"
	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] == num {
			cnt++
		} else {
			num = a[i]
			cnt = 1
		}
		if cnt == 3 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}