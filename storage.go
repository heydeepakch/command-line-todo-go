package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const todoFile = "todos.json"


func GetTodoFilePath() string {
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return todoFile
	}
	return filepath.Join(homeDir, ".todos.json")
}

func LoadTodos() (*TodoList, error) {
	filePath := GetTodoFilePath()
	
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &TodoList{Todos: []Todo{}}, nil
	}
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	
	var todoList TodoList
	if err := json.Unmarshal(data, &todoList); err != nil {
		return nil, err
	}
	
	return &todoList, nil
}


func SaveTodos(todoList *TodoList) error {
	filePath := GetTodoFilePath()
		
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return err
	}
	
	
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	
	return nil
}