package main

import (
	"cli-todo/internal/tools"
	"fmt"
)

func main() {
	fmt.Println("Starting todo cli")

	// var n = flag.Int("n", 123, "help message for flag n")

	// fmt.Printf("flag value: %v\r\n", *n)
	// flag.Parse()
	// fmt.Printf("flag value: %v\r\n", *n)

	db, err := tools.NewDatabase()
	if err != nil {
		fmt.Println("db is null")
	}

	// (*db).AddTodo(tools.Todo{Id: 1, Content: "testing"})
	// (*db).GetListOfTodos()
	// (*db).DeleteTodo(5)
	// (*db).AddTodo(tools.Todo{Id: 1, Content: "testing add after delete"})
	(*db).ChangeTodoCompleteStatus(1)
}
