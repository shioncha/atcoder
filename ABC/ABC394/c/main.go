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
	s := sc.Text()
	
	var builder strings.Builder
	builder.Grow(len(s))
	wcnt := 0
	
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'W':
			builder.WriteByte('W')
			wcnt++
		case 'A':
			if wcnt == 0 {
				builder.WriteByte('A')
			} else  {
				currentLen := builder.Len()
				tmp := builder.String()
				builder.Reset()
				builder.WriteString(tmp[:currentLen-wcnt])
				builder.WriteByte('A')
				builder.WriteString(strings.Repeat("C", wcnt))
				wcnt = 0
			}
		default:
			wcnt = 0
			builder.WriteByte(s[i])
		}
	}
	
	fmt.Println(builder.String())
}
