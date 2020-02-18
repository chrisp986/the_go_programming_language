package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) > 0 {
		arg := strings.Join(arg, " ")
		convert(arg)
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		convert(text)
	}
}

// TODO convert input and output it
func convert(command string) {
	fmt.Println(command)
}
