package main

import (
	"fmt"
	"strings"
	"szelvarix/greetings"
)

var name string

func main() {
	fmt.Print("Welcome adventurers, please give your name: ")
	fmt.Scanln(&name)
	nameToUpperCase := strings.ToUpper(name)
	greet := greetings.Hello(nameToUpperCase)
	fmt.Println(greet)
}
