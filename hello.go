package main

import (
	"fmt"

	"Desktop/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
