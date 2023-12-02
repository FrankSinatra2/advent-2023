package main

import (
	problem2 "advent/problem-2"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Expected 2 args")
		return
	}

	args := os.Args[1:]

	problems := map[string]func(){
		"problem-2": problem2.Run,
	}

	if problems[args[0]] != nil {
		problems[args[0]]()
	} else {
		fmt.Printf("Unkown input: %v\n", args[0])
	}
}
