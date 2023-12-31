package main

import (
	problem2 "advent/problem-2"
	problem3 "advent/problem-3"
	problem4 "advent/problem-4"
	problem5 "advent/problem-5"
	problem6 "advent/problem-6"
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
		"problem-3": problem3.Run,
		"problem-4": problem4.Run,
		"problem-5": problem5.Run,
		"problem-6": problem6.Run,
	}

	if problems[args[0]] != nil {
		problems[args[0]]()
	} else {
		fmt.Printf("Unkown input: %v\n", args[0])
	}
}
