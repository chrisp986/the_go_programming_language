package main

import (
	"bufio"
	"fmt"
	"os"
)

//Dup1 prints the text of each line that appears more than once in the standard input, preceded by its count
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func main() {
	Dup1()
}
