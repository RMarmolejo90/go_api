package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RMarmolejo90/go_api/database"
	"github.com/RMarmolejo90/go_api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder handles the creation of a new order in the database
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		slog.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&order); result.Error != nil {
		slog.Error("Failed to save order to the database", slog.String("error", result.Error.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetOrder retrieves an order from the database by ID
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Order not found", slog.String("id", id))
			c.JSON(http.StatusNotFound, gin.H{"error": "Order Not Found"})
		} else {
			slog.Error("Failed to retrieve order from the database", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetAllOrders retrieves all orders from the database
func GetAllOrders(c *gin.Context) {
	var orders []models.Order

	if err := database.DB.Find(&orders).Error; err != nil {
		slog.Error("Failed to retrieve orders from the database", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrder updates an order's information in the database by ID
func UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid order ID", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Order not found", slog.String("id", idStr))
			c.JSON(http.StatusNotFound, gin.H{"error": "Order Not Found"})
		} else {
			slog.Error("Failed to retrieve order from the database", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		slog.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Model(&order).Updates(updatedOrder); result.Error != nil {
		slog.Error("Failed to update order in the database", slog.String("error", result.Error.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	slog.Info("Successfully updated order", slog.Any("order", order))
	c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order from the database by ID
func DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid order ID", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Delete(&models.Order{}, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			slog.Info("Order not found", slog.String("id", idStr))
			c.JSON(http.StatusNotFound, gin.H{"error": "Order Not Found"})
		} else {
			slog.Error("Failed to delete order from the database", slog.String("error", result.Error.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	slog.Info("Successfully deleted order", slog.String("id", idStr))
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
}
