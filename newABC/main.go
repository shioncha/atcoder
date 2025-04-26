// Create directories and copy template files for ABC.

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	srcFile := "./Template/main.go"

	for i := 'a'; i < 'd'; i++ {
		dirName := fmt.Sprintf("./ABC/ABC%d/%c/", n, i)
		os.MkdirAll(dirName, 0755)
		copyFile(srcFile, dirName+"main.go")
	}

	fmt.Println("Create successfully")
}

func copyFile(srcFile, dstFile string) {
	src, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if  err != nil {
		panic(err)
	}
}
