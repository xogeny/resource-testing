package todo

type TodoItem struct {
	title     string
	completed bool
}

func (i TodoItem) Title() string {
	return i.title
}

func (i TodoItem) Completed() bool {
	return i.completed
}

func MakeItem(title string, completed bool) TodoItem {
	return TodoItem{
		title:     title,
		completed: completed,
	}
}
