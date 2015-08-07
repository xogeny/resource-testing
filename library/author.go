package library

import (
	"strings"
)

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

func (a Author) Matches(name string) bool {
	q := strings.ToLower(name)
	if strings.ToLower(a.firstName) == q {
		return true
	}
	for _, n := range a.middleNames {
		if strings.ToLower(n) == q {
			return true
		}
	}
	if strings.ToLower(a.lastName) == q {
		return true
	}
	return false
}

func MakeAuthor(firstName string, lastName string, middleNames ...string) Author {
	return Author{
		firstName:   firstName,
		middleNames: middleNames,
		lastName:    lastName,
	}
}
