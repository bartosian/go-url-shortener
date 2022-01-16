package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// func main() {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please enter your name:")
// 	input, err := inputReader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("There were errors reading, finishing program!")
// 		return
// 	}
// 	fmt.Printf("Your name is %s", input)

// 	switch input {
// 	case "Philip\n":
// 		fmt.Println("Welcome Philip")
// 	default:
// 		fmt.Printf("Who are you?")
// 	}
// }

// var nrchars, nrwords, nrlines int

// func main() {
// 	nrchars, nrlines, nrwords = 0, 0, 0
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please provide your text: ")

// 	for {
// 		input, err := inputReader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("There were errors reading text!")
// 			return
// 		}

// 		if input == "S\n" {
// 			fmt.Println("These are your counts: ")
// 			fmt.Println("Number of chars: %d\n", nrchars)
// 			fmt.Println("Number of words: %d\n", nrwords)
// 			fmt.Println("Number of lines: %d\n", nrlines)
// 			os.Exit(0)
// 		}
// 		Counters(input)
// 	}
// }

// func Counters(input string) {
// 	nrchars += len(input) - 1
// 	nrwords += len(strings.Fields(input))
// 	nrlines++
// }

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile.\n" +
			"Try open it and refactor.\n")
		return
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}

		fmt.Printf("The input was: %s", inputString)
	}
}
