package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a [10]string
	for i := 0; i < len(a); i++ {
		a[i] = strconv.Itoa(i)
	}
	fmt.Println(a)
}
