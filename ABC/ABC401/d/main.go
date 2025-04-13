package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Scanf("%s\n", &s)

	// dot_cnt := strings.Count(s, ".")
	o_cnt := strings.Count(s, "o")
	question_cnt := strings.Count(s, "?")

	if question_cnt == len(s) || (o_cnt == k && question_cnt == 0) {
		fmt.Print(s)
		return
	}
	ss := strings.Split(s, "")

	for i := strings.Index(s, "?"); i < len(s); i++ {
		if s[i] != '?' {
			continue
		}
		if i == 0 {
			if ss[i+1] == "o" {
				ss[i] = "."
			}
			continue
		}
		if len(s) < i + 2 {
			if ss[i-1] == "o" {
				ss[i] = "."
				continue
			}
			if s[i-1] == '.' && o_cnt == k {
				ss[i] = "o"
				o_cnt++
				continue
			}
		} else {
			if s[i-1] == 'o' && s[i+1] == '?' {
				ss[i] = "."
				continue
			}
			if s[i-1] == '.' && s[i+1] == '?' {
				ss[i] = "?"
				continue
			}
			if s[i-1] == '?' && s[i+1] == 'o' {
				ss[i-1] = "o"
				ss[i] = "."
				o_cnt++
				continue
			}
			if s[i-1] == '?' && s[i+1] == '.' {
				ss[i-1] = "."
				ss[i] = "o"
				continue
			}
		}
	}
	for _, v := range ss {
		fmt.Print(v)
	}
}