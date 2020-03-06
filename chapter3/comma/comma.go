package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBytes(s string) string {

	var buf bytes.Buffer
	for i, v := range s {

		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	return buf.String()
}

func main() {
	//fmt.Println(comma("123446"))
	fmt.Println(commaBytes("asb123456"))
}
