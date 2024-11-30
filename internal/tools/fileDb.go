package tools

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type fileDb struct {
	Index int //highest id index in the list to keep track for auto increments
}

const FILE_NAME string = "fileDb.txt"

func (db *fileDb) SetupDb() error {
	var file *os.File
	fi, err := os.Open(FILE_NAME)

	file = fi

	if os.IsNotExist(err) {
		fi, err := os.Create(FILE_NAME)
		if err != nil {
			fmt.Println("failed to create file db")
			return errors.New("failed to create file db")
		} else {
			db.Index = 0
			fmt.Println("Successfully created new file db")
		}
		file = fi
	} else {
		highest_index := 0
		sc := bufio.NewScanner(file)
		for sc.Scan() {
			line := sc.Text()
			todo, err := convertLineToTodo(line)
			if err != nil {
				return errors.New("failed to convert line to todo")
			}

			if todo.Id >= highest_index {
				highest_index = todo.Id
			}
		}
		db.Index = highest_index
	}

	file.Close()
	return nil
}

func (db *fileDb) GetListOfTodos() (map[int]Todo, error) {
	todos := make(map[int]Todo)

	fi, err := os.Open(FILE_NAME)
	if err != nil {
		return nil, errors.New("failed to open filedb")
	}

	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		line := sc.Text()
		t, err := convertLineToTodo(line)
		if err != nil {
			fmt.Println("failed to parse todo")
		}

		todos[t.Id] = t
	}

	fi.Close()
	return todos, nil
}

func (db *fileDb) DeleteTodo(id int) error {
	// read all the lines except the one with the id
	// write them all into the file and update db Index if necessary
	fi, err := os.Open(FILE_NAME)
	if err != nil {
		return errors.New("failed to open filedb")
	}

	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		line := sc.Text()
		t, err := convertLineToTodo(line)
		if err != nil {
			fmt.Println("failed to parse todo")
		}

		if t.Id == id {
			// delete the line and close the file
			return nil
		}
	}

	fi.Close()
	return fmt.Errorf("todo with id:%v not found", id)
}

func (db *fileDb) PatchTodo(id int) error {
	return errors.New("not implemented")
}

func (db *fileDb) AddTodo(todo Todo) error {
	// serialize todo as json
	// store line of <todoId>,<json>
	todo.Id = db.Index + 1
	j, err := json.Marshal(todo)
	if err != nil {
		return errors.New("failed to convert todo to json")
	}

	fi, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("failed to open filedb")
	}

	data := fmt.Sprintf("%d,%s\r\n", todo.Id, j)

	fi.WriteString(data)
	fi.Close()

	return nil
}

func convertLineToTodo(line string) (Todo, error) {
	line_arr := strings.Split(line, ",")
	todo_string := strings.Join(line_arr[1:], ",")
	var todo Todo
	json.Unmarshal([]byte(todo_string), &todo)

	return todo, nil
}
