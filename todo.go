package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}

func (tl *TodoList) AddTodo(title string) {
	newID := 1
	if len(tl.Todos) > 0 {
		for _, todo := range tl.Todos {
			if todo.ID >= newID {
				newID = todo.ID + 1
			}
		}
	}

	newTodo := Todo{
		ID:        newID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tl.Todos = append(tl.Todos, newTodo)
	fmt.Printf("Added todo #%d: %s\n", newTodo.ID, newTodo.Title)
}

func (tl *TodoList) CompleteTodo(id int)error{
	for i := range tl.Todos{
		if tl.Todos[i].ID ==id{
			tl.Todos[i].Completed = true
			now := time.Now()
			tl.Todos[i].CompletedAt = &now
			fmt.Printf("Completed todo #%d: %s\n", id, tl.Todos[i].Title)
			return nil
		}
	}
	return fmt.Errorf("todo with id %d not found", id)
}

func (tl *TodoList) DeleteTodo(id int)error{
	for i, todo := range tl.Todos{
		if todo.ID == id{
			tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
			fmt.Printf("Deleted todo #%d: %s\n", id, todo.Title)
			return nil
		}
	}
	return fmt.Errorf("todo with id %d not found", id)
	
}

func (tl *TodoList) ListTodos() {
	if len(tl.Todos) == 0 {
		fmt.Println("No todos yet! Add one with: todo add \"Your task\"")
		return
	}
	
	fmt.Println("\n📋 Your Todos:")
	fmt.Println("================")
	for _, todo := range tl.Todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s #%d: %s\n", status, todo.ID, todo.Title)
	}
	fmt.Println()
}
