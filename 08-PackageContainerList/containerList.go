package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Wahyu")
	data.PushBack("Sanjaya")
	data.PushBack("Alejandro")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Wahyu


	next := head.Next()
	fmt.Println(next.Value) // Sanjaya

	next = next.Next()
	fmt.Println(next.Value) // Alejandro

	for e := data.Front(); e!= nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}
 