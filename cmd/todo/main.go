package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pasDamola/todo-cli"
)

// Hardcoding the file name
const todoFileName = ".todo.json"


func main() {

	l := &todo.List{}
	
	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch  {
		// For no extra arguments, print the list
	case len(os.Args) == 1:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
		

	default:
		// Concatenate all arguments with a spcae
		item := strings.Join(os.Args[1:], " ")

		l.Add(item)
	}
}

