package handlers

import (
	"net/http"
	"strconv"

	"github.com/RMarmolejo90/go_api/api/database"
	"github.com/RMarmolejo90/go_api/api/models"
	"github.com/gin-gonic/gin"
)

func CreateCandle(c *gin.Context) {
	// bind json data to the model
	var candle models.Candle
	if err := c.BindJSON(&candle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.DB.Create(&candle); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, candle)
}

func GetCandle(c *gin.Context) {
	id := c.Param("id")
	var candle models.Candle

	if err := database.DB.First(&candle, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		return
	}
	c.JSON(http.StatusOK, candle)
}

func UpdateCandle(c *gin.Context) {

}
func DeleteCandle(c *gin.Context) {

}
