package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, finishing program!")
		return
	}
	fmt.Printf("Your name is %s", input)

	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip")
	default:
		fmt.Printf("Who are you?")
	}
}
