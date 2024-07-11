package models

import "time"

type Status string

const (
	NewOrder         Status = "newOrder"
	InProduction     Status = "inProduction"
	AwaitingShipment Status = "awaitingShipment"
	Shipped          Status = "shipped"
	Complete         Status = "complete"
	Other            Status = "other"
)

type Order struct {
	ID           uint      `gorm:"primary key"`       // Primary Key - Order Identifier
	User         User      `gorm:"foreignKey:UserID"` // Associated to this customer
	UserID       uint      // Customer ID foreign ID
	Products     []Candle  `gorm:"many2many:order_products" json:"products"`     // one to many relationship
	TotalPrice   float32   `json:"total"`                                        // total cost for the order in US dollars
	Tax          float32   `json:"tax"`                                          // tax amount
	ShippingCost float32   `json:"shipping"`                                     // shipping cost in US dollars
	Status       Status    `json:"status" gorm:"type:string;default:'newOrder'"` // order status - see type 'Status' at the top of this file
	CreatedAt    time.Time // time order was placed
	UpdatedAt    time.Time // time order was last updated
}
