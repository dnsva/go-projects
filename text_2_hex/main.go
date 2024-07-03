package main

import (
	"fmt"
	"log"

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

	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	gc.StartColor()     //enable color
	stdscr.Keypad(true) //enable keypad

	//gc.StdScr().Printf("Hello, Wofffrld!")
	gc.StdScr().Refresh()

	//--------------------------------------------------------------------------------------------
	//menu
	menu_options := []string{"1. Convert text to binary and hex", "2. Convert binary to text", "3. Convert hex to text", "4. Change max string size", "5. Exit"}

	var menu_choice int = 0

	for {

		// Hide cursor before displaying menu
		gc.Cursor(0)

		stdscr.MovePrint(0, 0, "Choose an option: ")
		for i, option := range menu_options {
			if i == menu_choice {
				stdscr.AttrOn(gc.A_STANDOUT)
			}
			stdscr.MovePrint(i+1, 0, option)
			stdscr.AttrOff(gc.A_STANDOUT)
			stdscr.Refresh() //so that everything is displayed
		}
		c := stdscr.GetChar() //get key
		if c == gc.KEY_DOWN {
			menu_choice = (menu_choice + 1) % 5
		} else if c == gc.KEY_UP {
			menu_choice = (menu_choice + 4) % 5
		} else if c == 10 {

			stdscr.Clear() //clear the screen
			gc.Cursor(1)   //show cursor

			if menu_choice == 0 {
				//text to binary and hex
				stdscr.Refresh()
			} else if menu_choice == 1 {
				stdscr.Refresh()
				//binary to text
			} else if menu_choice == 2 {
				stdscr.Refresh()
				//hex to text
			} else if menu_choice == 3 {
				stdscr.Refresh()
				//change max string size
			} else {
				//exit
				stdscr.Refresh()
				return
			}
		}
	}

	//char mesg[]="Enter a string: ";		/* message to be appeared on the screen */

	msg := "Enter a string (max size - 30): "
	//row, col := stdscr.MaxYX()
	//row, col = (row/2)-1, (col-len(msg))/2

	row, col := 0, 0
	stdscr.MovePrint(row, col, msg)

	var str string
	str, err = stdscr.GetString(30) //max 30 chars
	if err != nil {
		stdscr.MovePrint(row+1, col, "GetString Error:", err)
	} else {
		stdscr.MovePrintf(row+1, col, "You entered: %s", str)
		//Display binary and hex
		stdscr.MovePrintf(row+2, col, "Binary: %s", strToBinary(str))
		stdscr.MovePrintf(row+3, col, "Hex: %s", strToHex(str))
	}

	stdscr.Refresh()

	stdscr.GetChar()

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

		reader := bufio.NewReader(os.Stdin) //reader setup
		text, _ := reader.ReadString('\n')  //read the line from the user
		text = strings.TrimSpace(text)      //remove whitespace

		fmt.Println("Binary: ", strToBinary(text))
		fmt.Println("Hex: ", strToHex(text))

		fmt.Println("hello")*/

}
