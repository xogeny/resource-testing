package gotodo

type ID string

type TodoItem struct {
	title     string
	completed bool
}

func MakeItem(title string, completed bool) TodoItem {
	return TodoItem{
		title:     title,
		completed: completed,
	}
}
