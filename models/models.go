// models package contains definitions of data structures used in the application.
package models

// Stock struct represents the structure of a stock, with corresponding JSON field tags.
type Stock struct {
	StockID  int    `json:"stock_id"`  // Unique identifier for the stock
	Name     string `json:"name"`      // Name of the stock
	Price    int64  `json:"price"`     // Price of the stock in cents to avoid floating point imprecision
	Company  string `json:"company"`   // Company to which the stock belongs
}
