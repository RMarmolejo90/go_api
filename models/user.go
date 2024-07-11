package models

type User struct {
	ID        uint    `gorm:"primary_key"`                      // customer identifier
	FirstName string  `json:"firstName"`                        // customer first name
	LastName  string  `json:"lastName"`                         // customer last name
	Email     string  `json:"email"`                            // customer email
	Orders    []Order `json:"orders" gorm:"foreign_key:UserID"` // all order associated with the user
}
