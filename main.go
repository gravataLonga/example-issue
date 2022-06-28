package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := readFile()
	for {

		/**
		If enter 5, it split string correctly
		but if i enter "\n" don't read that.
		*/
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Bytes()
		fmt.Println(line)

		fmt.Println("result >>> ", bytes.Split(input, line))
	}
}

func readFile() []byte {
	b, r := ioutil.ReadFile("input.txt")
	if r != nil {
		fmt.Println(r)
		panic("Damn! ")
	}
	return b
}
