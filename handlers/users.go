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

func CreateUser(c *gin.Context) {
	// add user to the database
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// return the user info
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := database.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	idStr := c.Param("id")

	//convert id to integer for Gorm Query
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
	log.Printf("updated user %+v", user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {

	// convert id to integer
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// delete the user by id
	if result := database.DB.Delete(&models.User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully Deleted"})

}
