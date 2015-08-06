package gotodo

import (
	"fmt"
	uuid4 "github.com/nathanwinther/go-uuid4"
)

type TodoDatabase struct {
	open  bool
	items map[ID]TodoItem
}

func (db *TodoDatabase) AddTodo(title string, completed bool) (ID, error) {
	uuid, err := uuid4.New()
	if err != nil {
		return ID(""), err
	}

	id := ID(uuid)

	item := MakeItem(title, completed)
	db.items[id] = item

	return id, nil
}

func (db *TodoDatabase) Close() error {
	db.open = false
	return nil
}

func (db *TodoDatabase) Remove() error {
	if db.open {
		return fmt.Errorf("Cannot remove database while open")
	}
	return nil
}

func NewDB(name string) *TodoDatabase {
	ret := TodoDatabase{
		open:  true,
		items: map[ID]TodoItem{},
	}
	return &ret
}
