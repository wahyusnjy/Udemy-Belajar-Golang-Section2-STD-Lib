package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "password", "Put your database password")
	port := flag.Int("port", 3306, "Put your database port")


	flag.Parse()

	fmt.Println(*host, *username, *password, *port)
}
