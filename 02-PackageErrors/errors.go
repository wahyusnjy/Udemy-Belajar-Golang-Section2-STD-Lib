package main

import (
	"errors"
	"fmt"
)

var(
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found")
)

func GetByID(id string) (error) {
	if id == "" {
		return ValidationError
	}
	if id != "eko" {
		return NotFoundError
	}
	return nil
}
func main() {
	err := GetByID("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		}else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	} 
}
