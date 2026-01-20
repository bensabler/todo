package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bensabler/todo/internal/todo"
)

// Hard coding the file name
const fileName = ".todo.json"

func main() {

	// Define an items list
	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of argumennts provided
	switch {
	// For no extra arguments, just print the list
	case len(os.Args) == 1:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	// Concatenate all provided arguments with a space and
	// add to the list as an item
	default:
		// Concatenate all arguments with a space
		item := strings.Join(os.Args[1:], " ")

		// Add the task
		l.Add(item)

		// Save the new list
		if err := l.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	}

}
