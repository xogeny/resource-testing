package gotodo

type TodoDatabase struct {
}

func (db *TodoDatabase) Close() error {
	return nil
}

func (db *TodoDatabase) Remove() error {
	return nil
}

func NewDB(name string) *TodoDatabase {
	ret := TodoDatabase{}
	return &ret
}
