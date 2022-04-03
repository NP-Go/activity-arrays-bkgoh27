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

	var data [7]string
	data[0] = "Operating System List"
	data[1] = "Windows XP"
	data[2] = "Linux 1.0"
	data[3] = "Raspbian Teensy"
	data[4] = "MacOS"
	data[5] = "IOS"
	data[6] = "Google Andriod"

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i], len(data[i]))
	}

}
