package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RMarmolejo90/go_api/api/database"
	"github.com/RMarmolejo90/go_api/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	idStr := c.Param("id")
	var candle models.Candle

	// convert id to an integer for gorm query
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// find candle by ID
	if result := database.DB.First(&candle, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// bind data
	var updatedCandle models.Candle
	if err := c.ShouldBindJSON(&updatedCandle); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the candle
	if result := database.DB.Save(&candle); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Succuss Response
	log.Printf("Successfully Updated %+v", candle)
	c.JSON(http.StatusOK, candle)
}

func DeleteCandle(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	// convert id to an integer for gorm query
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// delete candle
	if result := database.DB.Delete(&models.Candle{}, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Candle Not Found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// success response
	log.Print("Successfully Deleted")
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
}
