package greetings

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var foundPrefix string

func Hello(name string) string {
	for _, nameChar := range name {
		isFound := false
		if nameChar == ' ' {
			foundPrefix += " "
			continue
		}
		for _, alphabetChar := range alphabet {
			findChar(string(alphabetChar), nameChar, &isFound)

			if isFound {
				break
			}

		}
		if !isFound {
			return fmt.Sprintf("Sorry, %s!", name)
		}
	}
	return fmt.Sprintf("Greetings, %s!", name)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func findChar(char string, nameChar rune, isFound *bool) {
	clearScreen()
	currentChar := foundPrefix + char
	fmt.Println(currentChar)
	time.Sleep(100 * time.Millisecond)
	if char == string(nameChar) {
		foundPrefix += char
		*isFound = true
	}
}
