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
	commands, verbs := initVerbsCmds()
	state := &GameState{
		Inventory: make(map[string]bool),
		Knowledge: make(map[string]bool),
	}
	for {
		cmdOptions(state, verbs)
		printInputLine()
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

func initVerbsCmds() (map[string]func(*GameState), []string) {
	verbs := []string{"try to see", "pick pockets", "light lighter",
		"look at shelf", "look at table", "look at bed", "look at door"}
	commands := make(map[string]func(*GameState))
	commands[verbs[0]] = cmdLook
	commands[verbs[1]] = cmdPickPockets
	commands[verbs[2]] = cmdLightLighter
	return commands, verbs
}

func cmdOptions(state *GameState, verbs []string) {
	switch {
	case state.Knowledge["lighterInPockets"] && !state.Inventory["lidLighter"]:
		text := fmt.Sprintf("[%s] [%s]", verbs[1], verbs[2])
		textPacer(text)
	default:
		text := fmt.Sprintf("[%s] [%s]", verbs[0], verbs[1])
		textPacer(text)
	}
}

func printInputLine() {
	fmt.Println("___")
	time.Sleep(300 * time.Millisecond)
	fmt.Print(">")
	time.Sleep(200 * time.Millisecond)
	fmt.Print(">")
	time.Sleep(100 * time.Millisecond)
	fmt.Print(">")
}

func cleanInput(text string) string {
	textTrimed := strings.TrimSpace(text)
	textTrimedLower := strings.ToLower(textTrimed)
	return textTrimedLower
}
