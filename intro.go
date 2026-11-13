package main

import (
	"fmt"
	"strings"
	"time"
)

func intro() {
	time.Sleep(250 * time.Millisecond)
	fmt.Print("\n         .......        ")
	time.Sleep(325 * time.Millisecond)
	fmt.Print("\n       .--:.::::.       ")
	time.Sleep(325 * time.Millisecond)
	fmt.Print("\n     ...::......:-.    ")
	textPacerIntro("BIG")
	time.Sleep(250 * time.Millisecond)
	fmt.Print("\n   .::-:.::=+**=::..    ")
	textPacerIntro("CAVE")
	time.Sleep(225 * time.Millisecond)
	fmt.Print("\n  .-::.::=#%%%%%-::.    ")
	textPacerIntro("ADVENTURE")
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\n  :=: ::=%%%%%%%#---.   ")
	time.Sleep(325 * time.Millisecond)
	fmt.Print("\n ::.::.:*%%%%%%%#-:.:   ")
	time.Sleep(325 * time.Millisecond)
	fmt.Print("\n.....:-#%%%%%%%%%=-.=.  ")
	time.Sleep(325 * time.Millisecond)
	fmt.Print("\n  .::-+%%%%%%%%%%=:::.  ")
	textPacerIntro("presented by:")
	time.Sleep(25 * time.Millisecond)
	fmt.Print("\n   ..-:+%%%%%%%%*=-.... ")
	textPacerIntro("Paul Hoffmann")
	time.Sleep(0 * time.Millisecond)
	fmt.Print("\n ...................... ")
	time.Sleep(2 * time.Second)
	fmt.Print("\n")
	text := "You are standing in dark. Its not clear to you what happened. The air smells wet, but is silent. You blink. Still dark. You reach out with your hands, but there is nothing you can get hold of. What should you do?"
	textPacer(text)
	time.Sleep(1 * time.Second)
}
func textPacerIntro(text string) {
	var words []string
	words = strings.Split(text, " ")
	for i, word := range words {
		for _, char := range word {
			fmt.Print(string(char))
			time.Sleep(25 * time.Millisecond)
		}
		if i != len(words)-1 {
			fmt.Print(" ")
		}
	}
}
