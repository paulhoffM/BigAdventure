package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type GameState struct {
	Inventory   map[string]bool
	Knowledge   map[string]bool
	ActiveVerbs map[string]bool
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	state := &GameState{
		Inventory:   make(map[string]bool),
		Knowledge:   make(map[string]bool),
		ActiveVerbs: make(map[string]bool),
	}
	commands := initVerbsCmds(state)
	state.ActiveVerbs["try to see"] = true
	state.ActiveVerbs["pick pockets"] = true
	state.ActiveVerbs["use ticker"] = true
	for {
		cmdOptions(state)
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

func initVerbsCmds(*GameState) map[string]func(*GameState) {
	verbs := []string{"try to see", "pick pockets", "light lighter",
		"look at shelf", "look at table", "look at bed", "look at door", "use ticker"}
	commands := make(map[string]func(*GameState))
	commands[verbs[0]] = cmdLook
	commands[verbs[1]] = cmdPickPockets
	commands[verbs[2]] = cmdLightLighter
	commands[verbs[3]] = cmdLookShelf
	commands[verbs[4]] = cmdLookTable
	commands[verbs[5]] = cmdLookBed
	commands[verbs[6]] = cmdLookDoor
	commands[verbs[7]] = cmdUseTicker
	return commands
}

func cmdOptions(state *GameState) {
	time.Sleep(200 * time.Millisecond)
	for verb, _ := range state.ActiveVerbs {
		if state.ActiveVerbs[verb] == true {
			fmt.Printf("[%s]", verb)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func printInputLine() {
	fmt.Println("\n___")
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
