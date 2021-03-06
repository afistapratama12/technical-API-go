package controllers

import (
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OrderInput struct {
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
}

func GetAllOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var orders []models.Order

	db.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrderByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	idOrder := c.Param("orderID")

	var order models.Order
	if err := db.Where("order_id = ?", idOrder).First(&order).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "order ID tidak ditemukan"})
		return
	}

	db.Find(&order, "order_id = ?", idOrder)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func InsertNewOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// checking for data insert
	var orderInput OrderInput
	if err := c.ShouldBindJSON(&orderInput); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	order := models.Order{
		CustomerName: orderInput.CustomerName,
		OrderedAt:    orderInput.OrderedAt,
	}

	db.Create(&order)

	c.JSON(http.StatusOK, gin.H{
		"status": "success create order",
		"order":  order,
	})
}

func UpdateOrderByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	idOrder := c.Param("orderID")

	// checking for parameter orderID
	var order models.Order
	if err := db.Where("order_id = ?", idOrder).First(&order).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "order ID tidak ditemukan"})
		return
	}

	// checking for data input update
	var orderInput OrderInput
	if err := c.ShouldBindJSON(&orderInput); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&order).Update(orderInput)

	c.JSON(http.StatusOK, gin.H{
		"status": "success update order",
		"order":  order,
	})
}

func DeleteOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	idOrder := c.Param("orderID")

	// checking for parameter orderID
	var order models.Order
	if err := db.Where("orderID = ?", idOrder).First(&order).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "order ID tidak ditemukan"})
		return
	}

	db.Delete(&order, "order_id = ?", idOrder)

	c.JSON(http.StatusOK, gin.H{"status": "success delete order"})
}
