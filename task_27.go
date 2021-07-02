package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	sSlice := strings.Split(s, " ")
	reversedStr := ""
	for i := len(sSlice) - 1; i >= 0; i-- {
		reversedStr += sSlice[i]
		if i != 0 {
			reversedStr += " "
		}
	}

	fmt.Println(reversedStr)
}
