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
			FindChar(string(alphabetChar), nameChar, &isFound)

			if isFound {
				break
			}

		}
		if !isFound {
			return fmt.Sprintf("Sorry, %s!", name)
		}
	}

	role, skill, stat := Stats()
	card := Card(name, role, skill, stat)
	return card
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func FindChar(char string, nameChar rune, isFound *bool) {
	ClearScreen()
	currentChar := foundPrefix + char
	fmt.Println(currentChar)

	time.Sleep(50 * time.Millisecond)
	if char == string(nameChar) {
		foundPrefix += char
		*isFound = true
	}
}
