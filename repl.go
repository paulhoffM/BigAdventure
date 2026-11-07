package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := make(map[string]func())
	commands["look"] = cmdLook

	for {
		fmt.Print(">>>")
		scanner.Scan()
		cliInput := cleanInput(scanner.Text())
		commands[cliInput]()
	}
}

func cmdLook() {
	text := "You look up and around. Different shades of black"
	textPacer(text)
}

func cleanInput(text string) string {
	textTrimed := strings.TrimSpace(text)
	textTrimedLower := strings.ToLower(textTrimed)
	return textTrimedLower
}
