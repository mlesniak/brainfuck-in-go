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
	var memory [65536]int
	pointer := 0

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			fmt.Printf("%c", memory[pointer])
		case ',':
			panic("To be implemented")
		case '[':
			// Find the matching ] if the current memory value is zero, do nothing otherwise.
			if memory[pointer] != 0 {
				continue
			}
			bracketCounter := 0
			breakLoop := false
			for j := i + 1; !breakLoop && j < len(code); j++ {
				switch code[j] {
				case ']':
					if bracketCounter == 0 {
						i = j
						breakLoop = true
					} else {
						bracketCounter--
					}
				case '[':
					bracketCounter++
				}
			}
		case ']':
			// Find the matching [ if the current memory value is not zero, continue otherwise.
			if memory[pointer] == 0 {
				continue
			}
			bracketCounter := 0
			breakLoop := false
			for j := i - 1; !breakLoop && j >= 0; j-- {
				switch code[j] {
				case '[':
					if bracketCounter == 0 {
						i = j
						breakLoop = true
					} else {
						bracketCounter--
					}
				case ']':
					bracketCounter++
				}
			}
		default:
			// Ignore it.
		}
	}
}

func readFile() string {
	fileName := os.Args[1]
	bytes, _ := ioutil.ReadFile(fileName)
	code := string(bytes)
	return code
}
