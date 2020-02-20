package main

// pc[i] is the population count of 1
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&i)
		//fmt.Printf("%b\n", pc[i])
	}
}

// Initial PopCount
//goos: darwin
//goarch: amd64
//BenchmarkPopCount-4   	1000000000	         0.321 ns/op
//PASS
//ok  	_/Users/christian/Documents/code/golang/the_go_programming_language/chapter2/popcount	0.362s
// PopCount returns the population count (number of set bits) of x.
//func PopCount(x uint64) int {
//	return int(pc[byte(x>>(0*8))] +
//		pc[byte(x>>(1*8))] +
//		pc[byte(x>>(2*8))] +
//		pc[byte(x>>(3*8))] +
//		pc[byte(x>>(4*8))] +
//		pc[byte(x>>(5*8))] +
//		pc[byte(x>>(6*8))] +
//		pc[byte(x>>(7*8))])
//}

// Exercise 2.3
//goos: darwin
//goarch: amd64
//BenchmarkPopCount-4   	55090441	        21.7 ns/op
//PASS
//ok  	_/Users/christian/Documents/code/golang/the_go_programming_language/chapter2/popcount	1.223s
// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var unit byte
	for i := 0; i < 8; i++ {
		unit += pc[byte(x>>(i*8))]
	}
	return int(unit)
}

func main() {
	PopCount(1)
}
