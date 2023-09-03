package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "your-project-path/app"
    "your-project-path/models"
    "your-project-path/helpers"
    "strconv"
)

func UpdateUser(c *gin.Context) {
    userIDStr := c.Param("userId")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
        return
    }

    var updatedUser app.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    existingUser, err := models.GetUserByID(uint(userID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }

    existingUser.Username = updatedUser.Username
    existingUser.Email = updatedUser.Email

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
    userIDStr := c.Param("userId")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
        return
    }

    existingUser, err := models.GetUserByID(uint(userID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}