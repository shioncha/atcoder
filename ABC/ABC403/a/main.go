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
	a := make([]int, n)
	sc.Scan()
	tmp := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(tmp[i])
	}

	ans := 0
	for i := 0; i < n; i += 2 {
		ans += a[i]
	}
	fmt.Println(ans)
}
