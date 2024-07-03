package main

import (
	"fmt"
	"log"

	gc "github.com/gbin/goncurses"
)

func main() {
	fmt.Println("hello")

	_, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()
}
