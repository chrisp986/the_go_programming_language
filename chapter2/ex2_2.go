package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var value float64
	var rate float64
	if len(os.Args[1:]) > 0 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		unit := os.Args[3]
		value, _ = strconv.ParseFloat(arg1, 64)
		rate, _ = strconv.ParseFloat(arg2, 64)

		convert(value, rate, unit)
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Println("Please enter the value, rate and unit you want to convert!")
			scanner.Scan()
			in := scanner.Text()
			parts := strings.Split(in, ",")
			value, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				fmt.Println("ERROR:", err)
				fmt.Print("\nPlease enter a valid number: ")
			}
			rate, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Println("ERROR:", err)
				fmt.Print("\nPlease enter a valid number: ")
			}
			unit := parts[2]

			convert(value, rate, unit)
		}
	}
}

// TODO convert input and output it
func convert(value float64, rate float64, unit string) {
	convertValue := value * rate
	fmt.Printf("Converted Value: %.2f , Conversion rate: %.2f , Unit: %s\n", convertValue, rate, unit)
}
