package models

type StockStatus string

const (
	InStock     StockStatus = "inStock"
	MadeToOrder StockStatus = "madeToOrder"
	OutOfStock  StockStatus = "outOfStock"
)

type Candle struct {
	Name          string      `json:"name"`       // product name
	Size          float32     `json:"size"`       // candle size in ounces
	Scent         string      `json:"scent"`      // the name of the candle scent
	Color         string      `json:"color"`      // color of the candle
	Price         float32     `json:"price"`      // listed sale price in US dollars
	MfgCost       float32     `json:"cost"`       // cost to manufacture, this can be used for calculations such as profit margins
	Weight        float32     `json:"weight"`     // product weight in ounces
	ID            uint        `gorm:"primaryKey"` // product identifier - Primary Key
	StockStatus   StockStatus `json:"status"`     // see StockStatus constants
	ActiveListing bool        `json:"active"`     // product listing is Deactivated or active
	Inventory     uint8       `json:"inventory"`  // stock inventory
	ImageURL      string      `json:"imageURL"`   // url path to an image uploaded to the cloud
}
