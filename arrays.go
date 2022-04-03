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

	var OperatingSystemList [6]string
	OperatingSystemList[0] = "Windows XP"
	OperatingSystemList[1] = "Linux 1.0"
	OperatingSystemList[2] = "Raspbian Teensy"
	OperatingSystemList[3] = "MacOS"
	OperatingSystemList[4] = "IOS"
	OperatingSystemList[5] = "Google Andriod"

	for i := 0; i < len(OperatingSystemList); i++ {
		fmt.Println(OperatingSystemList[i], len(OperatingSystemList[i]))
	}

	fmt.Println("Update")
	OperatingSystemList[0] = "Windows 10"
	OperatingSystemList[1] = "Linux 16.04"
	OperatingSystemList[2] = "Raspbian Buster"
	fmt.Println(OperatingSystemList)

	var NewOperatingSystemList [10]string
	// copy old OperatingSystemList to new OperatingSystemList array
	for i := 0; i < len(OperatingSystemList); i++ {
		NewOperatingSystemList[i] = OperatingSystemList[i]
	}
	NewOperatingSystemList[6] = "Ubuntu"
	NewOperatingSystemList[7] = "MS-Dos"
	NewOperatingSystemList[8] = "Solaris"
	fmt.Println(NewOperatingSystemList)
}
