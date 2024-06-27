package models

type User struct {
	ID        uint    `gorm:"primary key"` // customer identifier
	firstName string  // customer first name
	lastName  string  // customer last name
	email     string  // customer email
	orders    []Order // all order associated with the user
}
