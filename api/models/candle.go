package models

type Candle struct {
	name            string  // product name
	size            uint8   // candle size in ounces
	scent           string  // the name of the candle scent
	color           string  // color of the candle
	price           float32 // price in US dollars
	weight          float32 // product weight in ounces
	ID              uint    // product identifier - Primary Key
	quantityInStock uint16  // the quantity in inventory
	soldOut         bool    // product is sold out
}
