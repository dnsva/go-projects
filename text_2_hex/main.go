package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	gc "github.com/gbin/goncurses"
)

//============================================================

func strToBinary(text string) string {
	var result string = ""
	for _, c := range text {
		result += fmt.Sprintf("%b ", c)
	}
	return result
}

//============================================================

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

//============================================================

func strToHex(text string) string {
	var result string = ""
	for _, c := range text {
		result += fmt.Sprintf("%X ", c) //caps hexadecimal
	}
	return result
}

//============================================================

func hexToStr(hex string) string {
	var result string = ""
	if len(hex)%2 != 0 {
		//add leading 0
		hex = "0" + hex
	}
	for i := 0; i < len(hex); i += 2 {
		h := hex[i : i+2]
		charCode, _ := strconv.ParseInt(h, 16, 64)
		result += fmt.Sprintf("%c", charCode)
	}
	return result
}

//============================================================

func main() {

	stdscr, err := gc.Init() //intialize ncurses
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	stdscr.Keypad(true) //enable keypad
	gc.Cursor(0)        //hide cursor
	gc.StdScr().Refresh()

	//-----------------------------------------------------------

	menu_options := []string{"1. Convert text to binary and hex",
		"2. Convert binary to text",
		"3. Convert hex to text",
		"4. Change max string size", "5. Exit"}
	var menu_choice int = 0
	var max_input_size int = 30

	for {

		stdscr.Clear() //Clear anything previous

		stdscr.MovePrint(0, 0, "Choose an option: ")
		for i, option := range menu_options {
			if i == menu_choice {
				stdscr.AttrOn(gc.A_STANDOUT)
			}
			stdscr.MovePrint(i+1, 0, option)
			stdscr.AttrOff(gc.A_STANDOUT)
			stdscr.Refresh() //so that everything is displayed
		}
		//....................................................

		c := stdscr.GetChar() //get key
		if c == gc.KEY_DOWN {
			menu_choice = (menu_choice + 1) % 5

		} else if c == gc.KEY_UP {
			menu_choice = (menu_choice + 4) % 5

		} else if c == 10 {

			stdscr.Clear() //clear the screen

			////////////////////////////////////////////////////////
			if menu_choice == 0 { //text to binary and hex
				var msg string = "Enter a string (max size - " + strconv.Itoa(max_input_size) + "): "
				row, col := 0, 0
				stdscr.MovePrint(0, 0, msg)
				var str string
				str, err = stdscr.GetString(30) //max 30 chars
				if err != nil {
					stdscr.MovePrint(row+1, col, "GetString Error:", err)
				} else {
					stdscr.MovePrintf(row+1, col, "You entered: %s", str)
					stdscr.MovePrintf(row+2, col, "Binary: %s", strToBinary(str))
					stdscr.MovePrintf(row+3, col, "Hex: %s", strToHex(str))
				}
				stdscr.Refresh()
				stdscr.GetChar()
				////////////////////////////////////////////////////////
			} else if menu_choice == 1 { //binary to text
				var msg string = "Enter a decimal string (NO SPACES): "
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
				str = strings.TrimSpace(str)
				stdscr.MovePrintf(row+1, col, "You entered: %s\n", str)
				stdscr.MovePrintf(row+2, col, "Text: %s", binaryToStr(str))
				stdscr.Refresh()
				gc.Echo(true)
				stdscr.GetChar() //exit
				////////////////////////////////////////////////////////
			} else if menu_choice == 2 { //hex to text
				var msg string = "Enter a hexadecimal string (NO SPACES, CAPS): "
				gc.Echo(false)
				row, col := 0, 0
				stdscr.MovePrint(0, 0, msg)
				var str string = ""
				for j := 0; j < max_input_size*8; j++ {
					curr_char := stdscr.GetChar()
					if curr_char == 10 {
						break
					} else if curr_char >= '0' && curr_char <= '9' || curr_char >= 'A' && curr_char <= 'F' {
						stdscr.MovePrint(row, col+j+46, string(curr_char))
					} else {
						j--
						continue
					}
					str += string(curr_char)
					if err != nil {
						stdscr.MovePrint(row+1, col, "GetString Error:", err)
					}
				}
				str = strings.TrimSpace(str)
				stdscr.MovePrintf(row+1, col, "You entered: %s\n", str)
				stdscr.MovePrintf(row+2, col, "Text: %s", hexToStr(str))
				stdscr.Refresh()
				gc.Echo(true)
				stdscr.GetChar() //exit
				////////////////////////////////////////////////////////
			} else if menu_choice == 3 { //change max string size
				stdscr.Clear()
				stdscr.MovePrint(0, 0, "Enter new max string size: ")
				input, _ := stdscr.GetString(100)
				input = strings.TrimSpace(input)
				max_input_size, err = strconv.Atoi(input)
				if err != nil {
					stdscr.MovePrint(1, 0, "GetInteger Error:", err)
				}
				stdscr.MovePrintf(1, 0, "Max string size set to: %d", max_input_size)
				stdscr.Refresh()
				stdscr.GetChar() //exit
				////////////////////////////////////////////////////////
			} else { //exit
				stdscr.Refresh()
				return
			}
		}
	}
}
