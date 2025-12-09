package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Load existing todos
	todoList, err := LoadTodos()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading todos: %v\n", err)
		os.Exit(1)
	}
	
	// Check if any command was provided
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	
	command := os.Args[1]
	
	// Handle different commands
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo title")
			fmt.Println("Usage: todo add \"Your task\"")
			os.Exit(1)
		}
		title := os.Args[2]
		todoList.AddTodo(title)
		
	case "list", "ls":
		todoList.ListTodos()
		return // Don't save for read-only operations
		
	case "complete", "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo ID")
			fmt.Println("Usage: todo complete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			os.Exit(1)
		}
		if err := todoList.CompleteTodo(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		
	case "delete", "rm":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo ID")
			fmt.Println("Usage: todo delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			os.Exit(1)
		}
		if err := todoList.DeleteTodo(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		
	case "help":
		printUsage()
		return
		
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
	
	// Save todos after modifications
	if err := SaveTodos(todoList); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving todos: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("📝 Todo CLI - Manage your tasks from the command line")
	fmt.Println("\nUsage:")
	fmt.Println("  todo add \"Task description\"    - Add a new todo")
	fmt.Println("  todo list                      - List all todos")
	fmt.Println("  todo complete <id>             - Mark todo as complete")
	fmt.Println("  todo delete <id>               - Delete a todo")
	fmt.Println("  todo help                      - Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  todo add \"Buy groceries\"")
	fmt.Println("  todo list")
	fmt.Println("  todo complete 1")
	fmt.Println("  todo delete 2")
}