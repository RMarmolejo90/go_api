package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/RMarmolejo90/go_api/database"
	"github.com/RMarmolejo90/go_api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCandle handles the creation of a new candle in the database
func CreateCandle(c *gin.Context) {
	var candle models.Candle
	if err := c.ShouldBindJSON(&candle); err != nil {
		slog.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&candle); result.Error != nil {
		slog.Error("Failed to save candle to the database", slog.String("error", result.Error.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, candle)
}

// GetCandle retrieves a candle from the database by ID
func GetCandle(c *gin.Context) {
	id := c.Param("id")
	var candle models.Candle

	if err := database.DB.First(&candle, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Candle not found", slog.String("id", id))
			c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		} else {
			slog.Error("Failed to retrieve candle from the database", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, candle)
}

// UpdateCandle updates a candle's information in the database by ID
func UpdateCandle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid candle ID", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var candle models.Candle
	if err := database.DB.First(&candle, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Candle not found", slog.String("id", idStr))
			c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		} else {
			slog.Error("Failed to retrieve candle from the database", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updatedCandle models.Candle
	if err := c.ShouldBindJSON(&updatedCandle); err != nil {
		slog.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Model(&candle).Updates(updatedCandle); result.Error != nil {
		slog.Error("Failed to update candle in the database", slog.String("error", result.Error.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	slog.Info("Successfully updated candle", slog.Any("candle", candle))
	c.JSON(http.StatusOK, candle)
}

// DeleteCandle deletes a candle from the database by ID
func DeleteCandle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid candle ID", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Delete(&models.Candle{}, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			slog.Info("Candle not found", slog.String("id", idStr))
			c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		} else {
			slog.Error("Failed to delete candle from the database", slog.String("error", result.Error.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	slog.Info("Successfully deleted candle", slog.String("id", idStr))
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
}
