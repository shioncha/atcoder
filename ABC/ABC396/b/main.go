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
	q, _ := strconv.Atoi(sc.Text())
	query1 := make([]string, q)
	query2 := make([]int, q)
	for i := 0; i < q; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		query1[i] = s[0]
		if s[0] == "1" {
			query2[i], _ = strconv.Atoi(s[1])
		}
	}

	slice := make([]int, 0)
	for i := 0; i < q; i++ {
		if query1[i] == "1" {
			slice = append(slice, int(query2[i]))
		} else {
			if len(slice) == 0 {
				fmt.Println(0)
				continue
			}
			fmt.Println(slice[len(slice)-1])
			slice = slice[:len(slice)-1]
		}
	}
}