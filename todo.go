package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	CreatedAt time.Time `json:"created_at"`
}

type Todos []Todo

func NewTodo() Todo {
	return Todo{CreatedAt: time.Now()}
}
