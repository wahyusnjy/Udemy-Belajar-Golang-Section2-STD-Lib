package main 

import (
	"fmt"
	"sort"
)

type User struct {
	Name    string
	Age     int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main () {
	users := UserSlice{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// Sort the users by age
	sort.Sort(users)

	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
	}
	fmt.Println("Sorted users by age")
	fmt.Println("This is a simple example of sorting a slice of structs in Go using the sort package.")
}