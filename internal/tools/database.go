package tools

import "fmt"

type Todo struct {
	Id      int
	Content string
}

type DatabaseInterface interface {
	SetupDb() error
	GetListOfTodos() (map[int]Todo, error)
	DeleteTodo(id int) error
	PatchTodo(id int) error
	AddTodo(todo Todo) error
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &fileDb{}

	err := db.SetupDb()
	if err != nil {
		return nil, err
	}

	return &db, nil
}

func AddTodo() {
	var db DatabaseInterface = &fileDb{}

	db.AddTodo(Todo{Id: 99, Content: "testing one two, three"})
}

func GetListOfTodos() {
	var db DatabaseInterface = &fileDb{}
	todos, err := db.GetListOfTodos()
	if err != nil {
		fmt.Println("failed to get todos")
	}

	fmt.Println(todos[1])
}
