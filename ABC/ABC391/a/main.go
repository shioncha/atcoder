package main

import (
	"fmt"
)

func main() {
  var input string
  fmt.Scanf("%s", &input)

	switch input {
	case "N":
		fmt.Println("S")
	case "E":
		fmt.Println("W")
	case "W":
		fmt.Println("E")
	case "S":
		fmt.Println("N")
	case "NE":
		fmt.Println("SW")
	case "NW":
		fmt.Println("SE")
	case "SE":
		fmt.Println("NW")
	case "SW":
		fmt.Println("NE")
	}
}
