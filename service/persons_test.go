package service

import (
	"fmt"
	"testing"
)

func TestFindPerson(t *testing.T) {
	person := FindPerson(1)
	fmt.Println(person)
}

func TestGetPerson(t *testing.T) {

	whereKey := map[string]string{
		"email": "tarou@sample.com",
		"age":   "20",
	}
	person := GetPerson(whereKey)
	fmt.Println(person)
}
