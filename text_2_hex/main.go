package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	gc "github.com/gbin/goncurses"
)

func strToBinary(text string) string {
	var result string = ""
	for _, c := range text {
		result += fmt.Sprintf("%b ", c)
	}
	return result
}

func binaryToStr(binary string) string {
	var result string = ""
	for i := 0; i < len(binary); i += 8 {
		b := binary[i : i+8]
		result += string(b)
	}
	return result
}

func strToHex(text string) string {
	var result string = ""
	for _, c := range text {
		result += fmt.Sprintf("%X ", c) //caps hexadecimal
	}
	return result
}

func hexToStr(hex string) string {
	var result string = ""
	for i := 0; i < len(hex); i += 2 {
		h := hex[i : i+2]
		result += string(h) + " "
	}
	return result
}

func main() {

	_, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	gc.StartColor() //enable color

	gc.StdScr().Printf("Hello, World!")
	gc.StdScr().Refresh()

	/*
			setlocale(LC_ALL, ""); //enable utf-8
		    initscr();             //initialize screen
		    cbreak();              //other ncurses initialization
		    noecho();              //other ncurses initialization
		    curs_set(0);           //set cursor pos
		    keypad(stdscr, TRUE);  //set to standard screen output
		    start_color();         //enable color


			refresh(); //send everything to screen


			endwin(); //end ncurses mode

			printw("W

			mvprintw(i+1,0,"%d. %s\n", i+1, options[i].c_str()); //print option
		            attroff(A_STANDOUT); //unhighlight
	*/

	reader := bufio.NewReader(os.Stdin) //reader setup
	text, _ := reader.ReadString('\n')  //read the line from the user
	text = strings.TrimSpace(text)      //remove whitespace

	fmt.Println("Binary: ", strToBinary(text))
	fmt.Println("Hex: ", strToHex(text))

	fmt.Println("hello")

}
