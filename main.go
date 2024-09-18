package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, del int
	fmt.Scan(&num, &del)
	var rezString string

	stringNum := strconv.Itoa(num)
	for _, v := range stringNum {
		tempDigit, _ := strconv.Atoi(string(v))
		if tempDigit != del {
			rezString += strconv.Itoa(tempDigit)
		}
	}
	fmt.Println(rezString)

}
