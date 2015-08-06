package gotodo

import (
	"fmt"
	uuid4 "github.com/nathanwinther/go-uuid4"
)

type TodoDatabase struct {
	open  bool
	items map[ID]TodoItem
}

func (db *TodoDatabase) Get(id ID) (string, bool, error) {
	item, exists := db.items[id]
	if !exists {
		return "", false, fmt.Errorf("No item with id %s", id)
	}
	return item.title, item.completed, nil
}

func (db *TodoDatabase) ListAll() []ID {
	ret := []ID{}
	for k, _ := range db.items {
		ret = append(ret, k)
	}
	return ret
}

func (db *TodoDatabase) ListActive() []ID {
	ret := []ID{}
	for k, i := range db.items {
		if !i.completed {
			ret = append(ret, k)
		}
	}
	return ret
}

func (db *TodoDatabase) ListCompleted() []ID {
	ret := []ID{}
	for k, i := range db.items {
		if i.completed {
			ret = append(ret, k)
		}
	}
	return ret
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
