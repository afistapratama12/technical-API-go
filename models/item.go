package models

type Item struct {
	LineItemID uint   `gorm:"PRIMARY_KEY; AUTO_INCREMENT; NOT_NULL" json:"lineItemID"`
	ItemCode   string `json:"itemCode"`
	ItemName   string `json:"itemName"`
	Quantity   uint   `json:"quantity"`
	Order      Order  `json:"order" binding:"required" gorm:"foreignkey:OrderID"`
	OrderID    uint   `json:"orderID"`
}
