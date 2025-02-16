package main

import (
	"fmt"
)

func main() {
	var a string
	var b string

	fmt.Scan(&a, &b)

	if a == "fine" && b == "fine" {
		fmt.Println(4)
	} else if a == "fine" && b == "sick" {
		fmt.Println(3)
	} else if a == "sick" && b == "fine" {
		fmt.Println(2)
	} else {
		fmt.Println(1)
	}
}