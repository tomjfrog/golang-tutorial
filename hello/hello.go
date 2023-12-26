package main

import (
	"fmt"
	"tomjfrog.com/greetings"
)

func main() {
	message := greetings.Hello("Tom")
	fmt.Println(message)
}
