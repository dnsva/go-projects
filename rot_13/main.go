package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getIntFromInput(reader *bufio.Reader, minNum int, maxNum int) int {
	input, err := reader.ReadString('\n')
	if err != nil {
		return -1
	}
	input = strings.TrimSpace(input)
	numChoice, err := strconv.Atoi(input)
	if err != nil || numChoice < minNum || numChoice > maxNum {
		return -1
	}
	return numChoice
}

func getWordInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil { //error check
		return "-1"
	}
	input = strings.ToLower(strings.TrimSpace(input)) //remove whitespace and force lc
	for i := 0; i < len(input); i++ {
		if input[i] == 32 { //space is OK
			continue
		}
		if input[i] < 65 || input[i] > 122 || (input[i] > 90 && input[i] < 97) { //check for alpha
			return "-1"
		}
	}
	return input
}

func main() {

	reader := bufio.NewReader(os.Stdin) //reader setup

	for i := 1; i <= 5; i++ { //display options in a table
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d) ROT %d\t", (j-1)*5+i, (j-1)*5+i)
		}
		fmt.Println()
	}
	fmt.Print("Choose an option: ")
	choice := getIntFromInput(reader, 1, 25)
	for choice == -1 { //error check
		fmt.Print("Try again: ")
		choice = getIntFromInput(reader, 1, 25)
	}

	fmt.Print("Do you want to encode or decode?\n[1] Encode\n[2] Decode\n> ")
	shiftType := getIntFromInput(reader, 1, 2)
	for shiftType == -1 { //error check
		fmt.Print("Try again: ")
		shiftType = getIntFromInput(reader, 1, 2)
	}

	fmt.Print("Enter text: ") //text to shift or unshift
	input := getWordInput(reader)
	for input == "-1" { //error check
		fmt.Print("Try again: ")
		input = getWordInput(reader)
	}

	if choice < 10 { //formatting
		fmt.Printf("Using: %s\n", input)
	} else {
		fmt.Printf("Using : %s\n", input)
	}

	var shifted string //final result. Either shifting left or right depending on encoding or decoding
	for i := 0; i < len(input); i++ {
		if input[i] == 32 { //space is OK
			shifted += " "
			continue
		}
		if shiftType == 1 {
			shifted += (string)((int(input[i])-97+choice)%26 + 97) //shift right
		} else {
			//here the +26 is to make sure modulus is positive
			shifted += (string)((int(input[i])-97-choice+26)%26 + 97) //unshift (go left)

		}
	}

	if shiftType == 1 {
		fmt.Printf("ROT %d: %s\n", choice, shifted) //result
	} else {
		fmt.Printf("Original: %s\n", shifted) //result
	}
}
