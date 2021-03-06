package models

type Item struct {
	LineItemID uint   `gorm:"PRIMARY_KEY; AUTO_INCREMENT; NOT_NULL" json:"lineItemID"`
	ItemCode   string `json:"itemCode"`
	ItemName   string `json:"itemName"`
	Quantity   uint   `json:"quantity"`
	OrderID    uint   `gorm:"FOREIGN_KEY:OrderID" json:"orderID"`
}
