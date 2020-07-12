package model

//Person メンバーのモデル
type Person struct {
	ID         uint `gorm:"primary_key"`
	PersonName string
	Email      string
	Age        int
}
