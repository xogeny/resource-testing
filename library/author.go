package library

type Author struct {
	firstName   string
	middleNames []string
	lastName    string
}

func (a Author) FullName() string {
	ret := a.firstName
	for _, n := range a.middleNames {
		ret = ret + " " + n
	}
	ret = ret + " " + a.lastName
	return ret
}

func MakeAuthor(firstName string, lastName string, middleNames ...string) Author {
	return Author{
		firstName:   firstName,
		middleNames: middleNames,
		lastName:    lastName,
	}
}
