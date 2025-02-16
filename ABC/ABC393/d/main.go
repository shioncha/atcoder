package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1e6)
	sc.Buffer(buf, len(buf))

	sc.Scan()
	sc.Scan()
	s := sc.Text()

	ans := 0
	start := strings.Index(s, "1")
	end := strings.LastIndex(s, "1")
	cnt := strings.Count(s, "1")
	middle := 0
	tmp := 0
	for i:=start; i<=end; i++ {
		if s[i] == '0' {
			continue
		}
		if tmp == cnt/2 {
			middle = i
		}
		tmp++
	}
	lp := middle - 1
	for i:=middle-1; i>=start; i-- {
		if s[i] == '1' {
			ans += lp - i
			lp--
			// fmt.Printf("ans: %d", ans)
		}
	}
	rp := middle + 1
	for i:=middle+1; i<=end; i++ {
		if s[i] == '1' {
			ans += i - rp
			rp++
			// fmt.Printf("ans: %d", ans)
		}
	}
	// fmt.Println(start, end, cnt, middle)
	fmt.Println(ans)
}