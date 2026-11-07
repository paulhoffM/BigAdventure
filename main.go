package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//caveClient := cave.NewClient(5*time.Second, 5*time.Minute)
	intro()
	startRepl()
}

func textPacer(text string) {
	var words []string
	words = strings.Split(text, " ")
	fmt.Print("\n")
	for i, word := range words {
		fmt.Print(word + " ")
		time.Sleep(150 * time.Millisecond)
		if i%15 == 0 && i != 0 {
			fmt.Print("\n")
		}
		if i == len(words)-1 {
			fmt.Print("\n")
		}
	}
}

func intro() {
	fmt.Print("------------------")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\nBIG")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" CAVE")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" ADVENTURE")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\n------------------")
	time.Sleep(1 * time.Second)
	fmt.Print("\npresented by:")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" Paul Hoffmann")
	time.Sleep(1 * time.Second)
	fmt.Print("\n")
	time.Sleep(1 * time.Second)
	text := "You are standing in dark. Its not clear to you what happened. The air smells wet, but is silent. You blink. Still dark. You reach out with your hands, but there is nothing you can get hold of. What should you do?"
	textPacer(text)
}
