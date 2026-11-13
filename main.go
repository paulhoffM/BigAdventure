package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//	intro()
	startRepl()
}

func textPacer(text string) {
	var words []string
	words = strings.Split(text, " ")
	fmt.Print("\n")
	for i, word := range words {
		for _, char := range word {
			fmt.Print(string(char))
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Print(" ")
		if i%15 == 0 && i != 0 {
			fmt.Print("\n")
		}
		if i == len(words)-1 {
			fmt.Print("\n")
		}
	}
}
