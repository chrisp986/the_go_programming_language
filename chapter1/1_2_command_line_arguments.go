package main

import (
	"fmt"
	"os"
	"strings"
)

//Cla returning os.Args
func Cla() string {
	// var s, sep string

	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// return s + " <- this is it"

	return strings.Join(os.Args[0:], " ")
}

func main() {
	fmt.Println(Cla())
	// start := time.Now()

	// secs := time.Since(start).Nanoseconds()
	// fmt.Println(secs)
}
