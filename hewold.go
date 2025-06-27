package main

import (
	"fmt"
	"hewold/greetings"
	"strings"
)

var name string

func main() {
	fmt.Print("Welcome adventurers, please give your name: ")
	fmt.Scanln(&name)
	nameToUpperCase := strings.ToUpper(name)
	greet := greetings.Hello(nameToUpperCase)
	fmt.Println(greet)
}
