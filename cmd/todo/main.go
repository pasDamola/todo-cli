package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pasDamola/todo-cli"
)

// Hardcoding the file name
const todoFileName = ".todo.json"


func main() {

	// Parsing command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	l := &todo.List{}
	
	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on provided flags
	switch  {
		// For no extra arguments, print the list
	case *list:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} 

	case *task != "":
		// Add the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)

		}		

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

