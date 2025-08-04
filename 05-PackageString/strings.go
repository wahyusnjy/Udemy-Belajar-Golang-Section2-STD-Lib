package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Wahyu Sanjaya", "Wahyu"))
	fmt.Println(strings.Split("Wahyu Sanjaya", " "))
	fmt.Println(strings.ToLower("Wahyu Sanjaya"))
	fmt.Println(strings.ToUpper("Wahyu Sanjaya"))
	fmt.Println(strings.Trim("   Wahyu Sanjaya   ", " "))
	fmt.Println(strings.ReplaceAll("Wahyu Sanjaya", "Wahyu", "Alejandro"))

}
