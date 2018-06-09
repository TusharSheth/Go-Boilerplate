package main

import "time"

// Todo : todo
type Todo struct {
	Id        int       `json:Id`
	Name      string    `json:"name"`
	Completed bool      `json:"completed`
	Due       time.Time `json:"due"`
}

// Todos : todos
type Todos []Todo
