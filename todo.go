package gotodo

import (
	"fmt"
	uuid4 "github.com/nathanwinther/go-uuid4"
)

// This is here just to summarize the interface of a TodoDatabase
type TodoInterface interface {
	Get(ID) (item TodoItem, err error)
	ListAll() []ID
	ListActive() []ID
	ListCompleted() []ID
	AddTodo(TodoItem) (ID, error)
	Edit(id ID, item TodoItem) error
	MarkCompleted(ID) error
	ClearCompleted(ID) error
	Remove(ID) error
	Close() error
}

// This is the actual TODO database type
type TodoDatabase struct {
	// Mapping of IDs to actual todo items
	items map[ID]TodoItem
}

func (db *TodoDatabase) Get(id ID) (TodoItem, error) {
	item, exists := db.items[id]
	if !exists {
		return item, fmt.Errorf("No item with id %s", id)
	}
	return item, nil
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

func (db *TodoDatabase) AddTodo(item TodoItem) (ID, error) {
	uuid, err := uuid4.New()
	if err != nil {
		return ID(""), err
	}

	id := ID(uuid)

	db.items[id] = item

	return id, nil
}

func (db *TodoDatabase) Edit(id ID, item TodoItem) error {
	_, exists := db.items[id]
	if !exists {
		return fmt.Errorf("No item with id %s", id)
	}
	db.items[id] = item
	return nil
}

func (db *TodoDatabase) MarkCompleted(id ID) error {
	item, exists := db.items[id]
	if !exists {
		return fmt.Errorf("No item with id %s", id)
	}
	item.completed = true
	db.items[id] = item
	return nil
}

func (db *TodoDatabase) ClearCompleted(id ID) error {
	item, exists := db.items[id]
	if !exists {
		return fmt.Errorf("No item with id %s", id)
	}
	item.completed = false
	db.items[id] = item
	return nil
}

func (db *TodoDatabase) Remove(id ID) error {
	_, exists := db.items[id]
	if exists {
		delete(db.items, id)
		return nil
	}
	return fmt.Errorf("No item with id %s", id)
}

func (db *TodoDatabase) Close() error {
	return nil
}

func NewDB(name string) *TodoDatabase {
	ret := TodoDatabase{
		items: map[ID]TodoItem{},
	}
	return &ret
}

var _ TodoInterface = (*TodoDatabase)(nil)
