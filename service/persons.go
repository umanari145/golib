package service

import (
	"golib/db"
	"golib/model"
)

//FindPerson メンバーの取得
func FindPerson(id int) model.Person {

	person := model.Person{}

	dbCon, _ := db.Connect()

	dbCon.Table("persons").Find(&person, id)

	return person
}

//GetPerson メンバーの取得
func GetPerson(whereKey map[string]string) model.Person {

	person := model.Person{}

	dbCon, _ := db.Connect()

	dbCon.Table("persons").Where(whereKey).Find(person)

	return person
}
