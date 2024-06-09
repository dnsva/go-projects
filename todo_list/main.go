package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

type Todo struct {
	Title    string
	Progress bool //if false it means not started
}

var progressIndexEnd int = -1 //the last index with an in progress element

func AddItem(list *[]Todo, item Todo) {

	var copy []Todo = *list

	if item.Progress {

		copy = append(copy, Todo{"placeholder", false})

		for i := len(copy) - 1; i > progressIndexEnd+1; i-- {
			copy[i] = copy[i-1]
		}

		copy[progressIndexEnd+1] = item
		progressIndexEnd++

		list = &copy

	} else {
		*list = append(*list, item)
	}

	WriteData(*list)
}

func SortList(list *[]Todo) {
	var copy []Todo = *list
	var sorted []Todo

	for i := 0; i < len(copy); i++ {

		if copy[i].Progress {
			sorted = append(sorted, copy[i])
		}

	}

	progressIndexEnd = len(sorted) - 1

	for i := 0; i < len(copy); i++ {

		if !copy[i].Progress {
			sorted = append(sorted, copy[i])
		}

	}

	list = &copy

	WriteData(*list)
}

func RemoveItem(list *[]Todo, index int) {

	var copy []Todo = *list

	if copy[index].Progress {
		progressIndexEnd--
	}

	for i := index; i < len(copy)-1; i++ {
		copy[i] = copy[i+1]
	}

	*list = copy[:len(copy)-1]

	WriteData(*list)

	//*list = copy
}

func WriteData(list []Todo) {

	f, _ := os.Create("data.txt")
	//defer f.Close()

	dataWrite := ""
	for i := 0; i < len(list); i++ {
		dataWrite += list[i].Title + " " + strconv.FormatBool(list[i].Progress) + "\n"
	}

	f.WriteString(dataWrite)
	f.Sync()
	f.Close()
}

func LoadData(listOk *[]Todo) {

	data, _ := os.ReadFile("data.txt")

	lines := strings.Split(string(data), "\n")
	//lines = lines[1:] //skip the first weird ""

	var todoList []Todo

	for _, line := range lines {

		//fmt.Println(line)

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		splittedString := strings.Split(line, " ")

		if len(splittedString) != 2 {
			fmt.Println("Invalid data in file")
		} else {
			todoList = append(todoList, Todo{splittedString[0], splittedString[1] == "true"})
		}
	}

	SortList(&todoList)
	*listOk = todoList

}

func returnCharString(c rune, mult int) string {

	var str string = ""

	for i := 0; i < mult; i++ {
		str += string(c)
	}

	return str
}

func main() {

	var list []Todo
	reader := bufio.NewReader(os.Stdin)

	LoadData(&list)

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	for {

		//fmt.Print("\033[0M") //reset

		fmt.Print("\033[38;5;122m") //teal
		fmt.Println(returnCharString('=', width/2-6), "TO DO LIST", returnCharString('=', width/2-6))
		fmt.Print("\033[0m") //reset

		fmt.Print("\033[38;5;117m") //blue
		fmt.Println("In Progress", returnCharString('-', width/2*2-12))
		fmt.Print("\033[0m") //reset

		if progressIndexEnd == -1 {
			fmt.Println("No items in progress")
		} else {
			for i := 0; i <= progressIndexEnd; i++ {
				fmt.Printf("%d. %s\n", i, list[i].Title)
			}
		}

		fmt.Print("\033[38;5;117m") //blue
		fmt.Println("Not Started", returnCharString('-', width/2*2-12))
		fmt.Print("\033[0m") //reset

		if len(list)-progressIndexEnd-1 <= 0 {
			fmt.Println("No items not started")
		} else {
			for i := progressIndexEnd + 1; i < len(list); i++ {
				fmt.Printf("%d. %s\n", i, list[i].Title)
			}
		}

		fmt.Println(returnCharString('=', width/2*2))
		fmt.Println("[A] - Add Item                ")
		fmt.Println("[R] - Remove Item             ")
		fmt.Println("[P] - Mark as in progress     ")
		fmt.Println("[Q] - Quit                    ")
		fmt.Print("> ")

		selection, _ := reader.ReadString('\n')
		selection = strings.Trim(selection, "\n")

		//fmt.Printf("selection is %t", strings.ToUpper(selection) == "A")
		//fmt.Print("hello>?")

		if strings.ToUpper(selection) == "A" {
			fmt.Print("Enter the title of the item: ")
			title, _ := reader.ReadString('\n')
			title = strings.Trim(title, "\n")
			AddItem(&list, Todo{title, false})
		} else if strings.ToUpper(selection) == "R" {

			if len(list) == 0 {
				fmt.Println("No items to remove")
			} else {
				fmt.Print("Enter task number to removed: ")
				indexStr, _ := reader.ReadString('\n')
				indexStr = strings.Trim(indexStr, "\n")
				index, err := strconv.Atoi(indexStr)
				for err != nil || index >= len(list) || index < 0 {
					fmt.Print("Invalid, try again: ")
					indexStr, _ := reader.ReadString('\n')
					indexStr = strings.Trim(indexStr, "\n")
					index, err = strconv.Atoi(indexStr)
				}
				//fmt.Printf("the length aws %d\n", len(list))
				RemoveItem(&list, index)
				SortList(&list)
			}
		} else if strings.ToUpper(selection) == "P" {

			if len(list) == 0 {
				fmt.Println("No items to mark as in progress")
			} else {
				fmt.Print("Enter progress task number: ")
				indexStr, _ := reader.ReadString('\n')
				indexStr = strings.Trim(indexStr, "\n")
				index, err := strconv.Atoi(indexStr)
				for err != nil || index > len(list) || index < 0 {
					fmt.Print("Invalid, try again: ")
					indexStr, _ := reader.ReadString('\n')
					indexStr = strings.Trim(indexStr, "\n")
					index, err = strconv.Atoi(indexStr)
				}
				list[index].Progress = true
				progressIndexEnd++
				SortList(&list)
			}
		} else if strings.ToUpper(selection) == "Q" {
			break
		} else {
			fmt.Println("Invalid selection")
		}
	}
}
