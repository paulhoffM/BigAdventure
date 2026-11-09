package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type GameState struct {
	Inventory map[string]bool
	Knowledge map[string]bool
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := make(map[string]func(*GameState))
	commands["try to see"] = cmdLook
	commands["pick pockets"] = cmdPickPockets
	commands["light lighter"] = cmdLightLighter
	state := &GameState{
		Inventory: make(map[string]bool),
		Knowledge: make(map[string]bool),
	}
	for {
		fmt.Println("___")
		time.Sleep(300 * time.Millisecond)
		fmt.Print(">")
		time.Sleep(200 * time.Millisecond)
		fmt.Print(">")
		time.Sleep(100 * time.Millisecond)
		fmt.Print(">")
		scanner.Scan()
		cliInput := cleanInput(scanner.Text())
		cmd, ok := commands[cliInput]
		if !ok {
			fmt.Println("You cant do that")
		} else {
			cmd(state)
		}

	}
}

func cleanInput(text string) string {
	textTrimed := strings.TrimSpace(text)
	textTrimedLower := strings.ToLower(textTrimed)
	return textTrimedLower
}
