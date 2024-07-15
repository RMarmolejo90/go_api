package handlers

import (
	"net/http"

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

	
func UpdateUser(c *gin.Context) {
	var user models.User
	idStr := c.Param("id")

	//convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user ID"})
		return
	}

	// Find user by ID
	if result := database.DB.First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// bind the json data to the model
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prevent changing of email and ID
	updatedUser.ID = user.ID
	updatedUser.Email = user.Email

	// Update the user data
	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName

	// save to the database
	if result := database.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// success
	log.Printf("Updated Candle %+v", candle)
	c.JSON(http.StatusOK, candle)
}

}
func DeleteCandle(c *gin.Context) {

}
