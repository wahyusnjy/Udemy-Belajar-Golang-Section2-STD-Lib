package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Boolean value:", boolean)

	int, err := strconv.Atoi("1230")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("int value:", int)


	binary := strconv.FormatInt(999, 2)
	fmt.Println("Binary :",binary)


	StringToInt := strconv.Itoa(999)
	fmt.Println("String to Int:", StringToInt)
}
