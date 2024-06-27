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
	ID           uint      `gorm:"primary key"` // Primary Key - Order Identifier
	customer     User      `gorm:"foreign key"` // foreign key - customer ID
	products     []Candle  // one to many relationship
	totalPrice   float32   // total cost for the order in US dollars
	tax          float32   // tax amount
	shippingCost float32   // shipping cost in US dollars
	status       Status    `gorm:"type:string;default:'newOrder'"` // order status - see type 'Status' at the top of this file
	createdAt    time.Time // time order was placed
	updatedAt    time.Time // time order was last updated
}
