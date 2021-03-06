package controllers

import (
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ItemInput struct {
	LineItemID uint   `json:"lineItemID"`
	ItemCode   string `json:"itemCode"`
	ItemName   string `json:"itemName"`
	Quantity   uint   `json:"quantity"`
	OrderID    uint   `json:"orderID"`
}

func GetAllItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var items []models.Item
	db.Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func GetItemByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// checking for params
	var item models.Item

	idItem := c.Param("lineItemID")

	println(idItem)

	if err := db.Where("line_item_id = ?", idItem).First(&item).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "item ID tidak ditemukan"})
		return

	}

	db.Find(&item, "line_item_id = ?", idItem)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

func InsertNewItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// checking for data insert
	var itemInput ItemInput
	if err := c.ShouldBindJSON(&itemInput); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	item := models.Item{
		ItemCode: itemInput.ItemCode,
		ItemName: itemInput.ItemName,
		Quantity: itemInput.Quantity,
		OrderID:  itemInput.OrderID,
	}

	db.Create(&item)

	c.JSON(http.StatusOK, gin.H{
		"status": "success create item",
		"item":   item,
	})
}

func DeleteItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	idItem := c.Param("lineItemID")

	// checking for params
	var item models.Item
	if err := db.Where("line_item_id = ?", idItem).First(&item).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "item ID tidak ditemukan"})
		return
	}

	// process delete
	db.Delete(&item, "line_item_id = ?", idItem)

	c.JSON(http.StatusOK, gin.H{"status": "success delete item"})
}
