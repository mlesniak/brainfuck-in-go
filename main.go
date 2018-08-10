package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	code := readFile()
	fmt.Println(code)
}

func readFile() string {
	fileName := os.Args[1]
	bytes, _ := ioutil.ReadFile(fileName)
	code := string(bytes)
	return code
}
