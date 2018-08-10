package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	code := readFile()
	parse(code)
}
func parse(code string) {
	for _, v := range code {
		fmt.Println(v)
	}
}

func readFile() string {
	fileName := os.Args[1]
	bytes, _ := ioutil.ReadFile(fileName)
	code := string(bytes)
	return code
}
