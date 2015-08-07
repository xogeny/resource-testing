package gotodo

type ID string

type todoItem struct {
	title     string
	completed bool
}

func makeItem(title string, completed bool) todoItem {
	return todoItem{
		title:     title,
		completed: completed,
	}
}
