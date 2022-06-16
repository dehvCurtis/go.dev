package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Robert")
	fmt.Println(message)
}
