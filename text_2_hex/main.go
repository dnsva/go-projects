package main

import (
	"fmt"
	"log"
	"strconv"
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

	if len(binary)%8 != 0 {
		//add leading 0s
		binary = strings.Repeat("0", 8-(len(binary)%8)) + binary
	}

	for i := 0; i < len(binary); i += 8 {
		b := binary[i : i+8]
		charCode, _ := strconv.ParseInt(b, 2, 64)
		result += fmt.Sprintf("%c", charCode)
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

	var max_input_size int = 30

	for {

		// Hide cursor before displaying menu
		gc.Cursor(0)

		//Clear anything previous
		stdscr.Clear()

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

			if menu_choice == 0 {
				//text to binary and hex
				var msg string = "Enter a string (max size - " + strconv.Itoa(max_input_size) + "): "

				row, col := 0, 0
				stdscr.MovePrint(0, 0, msg)

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

			} else if menu_choice == 1 {
				//binary to text
				var msg string = "Enter a binary string (NO SPACES): "

				gc.Echo(false)

				row, col := 0, 0
				stdscr.MovePrint(0, 0, msg)

				var str string = ""
				for j := 0; j < max_input_size*8; j++ {
					curr_char := stdscr.GetChar()
					if curr_char == 10 {
						break
					} else if curr_char == '1' || curr_char == '0' {
						stdscr.MovePrint(row, col+j+35, string(curr_char))
					} else {
						j--
						continue
					}
					str += string(curr_char)
					if err != nil {
						stdscr.MovePrint(row+1, col, "GetString Error:", err)
					}
				}
				//stdscr.Clear()
				stdscr.MovePrintf(row+1, col, "You entered: %s", str)
				//Display binary and hex
				//trim sapce:
				str = strings.TrimSpace(str)
				stdscr.MovePrintf(row+2, col, "Text: %s", binaryToStr(str))
				stdscr.Refresh()

				gc.Echo(true)

				stdscr.GetChar()

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

}
