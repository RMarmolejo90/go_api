package handlers

import (
	"net/http"

	"github.com/RMarmolejo90/go_api/api/database"
	"github.com/RMarmolejo90/go_api/api/models"
	"github.com/gin-gonic/gin"
)

func CreateCandle(c *gin.Context) {
	// bind json data to the model
	var candle models.Candle
	if err := c.ShouldBindJSON(&candle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// save candle to the database
	if result := database.DB.Create(&candle); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// return candle to client
	c.JSON(http.StatusOK, candle)
}

func GetCandle(c *gin.Context) {
	id := c.Param("id")
	var candle models.Candle

	// get candle from database
	if err := database.DB.First(&candle, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		return
	}
	// return candle
	c.JSON(http.StatusOK, candle)
}

func UpdateCandle(c *gin.Context) {
	id := c.Param("id")
	var candle models.Candle

	// bind data
	if err := c.ShouldBindJSON(&candle); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// update the candle
	if result := database.DB.update
}
func DeleteCandle(c *gin.Context) {

}
