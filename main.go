package main

import (
	"fmt"
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

// func main() {
// 	inputFile, inputError := os.Open("input.dat")
// 	if inputError != nil {
// 		fmt.Printf("An error occurred on opening the inputfile.\n" +
// 			"Try open it and refactor.\n")
// 		return
// 	}

// 	defer inputFile.Close()
// 	inputReader := bufio.NewReader(inputFile)

// 	for {
// 		inputString, readerError := inputReader.ReadString('\n')
// 		if readerError == io.EOF {
// 			return
// 		}

// 		fmt.Printf("The input was: %s", inputString)
// 	}
// }

// func main() {
// 	inputFile := "products.txt"
// 	buf, err := ioutil.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Printf(os.Stderr, "File Error: %s\n", err)
// 	}
// }

func main() {
	file, err := os.Open("products2.tx")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}

		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
