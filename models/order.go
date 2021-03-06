package models

type Order struct {
	OrderID      uint    `json:"orderID" gorm:"PRIMARY_KEY; AUTO_INCREMENT; NOT_NULL"`
	CustomerName string  `json:"customerName"`
	OrderedAt    string  `json:"orderedAt"`
	Items        *[]Item `json:"items"`
}
