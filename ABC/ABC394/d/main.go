package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1e6)
	sc.Buffer(buf, len(buf))

	sc.Scan()
	s := sc.Text()

	c1 := 0
	c2 := 0
	c3 := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			c1++
		case ')':
			c1--
		case '[':
			c2++
		case ']':
			c2--
		case '<':
			c3++
		case '>':
			c3--
		}
		if c1 < 0 || c2 < 0 || c3 < 0 {
			fmt.Println("No")
			return
		}
	}

	st := make([]int, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			st = append(st, 0)
		case '[':
			st = append(st, 1)
		case '<':
			st = append(st, 2)
		case ')':
			if len(st) == 0 || st[len(st)-1] != 0 {
				fmt.Println("No")
				return
			}
			st = st[:len(st)-1]
		case ']':
			if len(st) == 0 || st[len(st)-1] != 1 {
				fmt.Println("No")
				return
			}
			st = st[:len(st)-1]
		case '>':
			if len(st) == 0 || st[len(st)-1] != 2 {
				fmt.Println("No")
				return
			}
			st = st[:len(st)-1]
		}
	}
	if len(st) > 0 {
		fmt.Println("No")
		return
	}
	
	fmt.Println("Yes")
}
