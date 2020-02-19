package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var value float64
	var rate float64
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	unit := os.Args[3]
	value, _ = strconv.ParseFloat(arg1, 64)
	rate, _ = strconv.ParseFloat(arg2, 64)
	if value > 0 {
		//arg := strings.Join(arg, " ")
		convert(value, rate, unit)
	}
	//else {
	//	reader := bufio.NewReader(os.Stdin)
	//	fmt.Print("Enter text: ")
	//	text, _ := reader.ReadString('\n')
	//	convert(text)
	//}
}

// TODO convert input and output it
func convert(value float64, rate float64, unit string) {
	fmt.Printf("Value: %.2f , Conversion rate: %.2f , Unit: %s", value, rate, unit)
}
