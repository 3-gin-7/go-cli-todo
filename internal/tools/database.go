package tools

type Todo struct {
	Id      int
	Content string
	IsDone  bool
}

type DatabaseInterface interface {
	SetupDb() error
	GetListOfTodos() (map[int]Todo, error)
	DeleteTodo(id int) error
	ChangeTodoCompleteStatus(id int) error
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
