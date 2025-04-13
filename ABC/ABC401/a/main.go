package main

import (
	"fmt"
)

func main() {
  var s int
  fmt.Scanf("%d", &s)

  if 200 <= s && s <= 299 {
    fmt.Println("Success")
  } else {
    fmt.Println("Failure")
  }
}
