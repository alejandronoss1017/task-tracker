package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Expected at least one argument, got none")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		// TODO: Implement add task to the list
	case "list":
		// TODO: Implement list all the tasks logic
	case "remove":
		// TODO: Implement remove task logic
	case "update":
		// TODO: Implement update task logic
	}
}
