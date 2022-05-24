package controllers


import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/cloudguruab/rest-api-golang/controllers"
)


func FindBooks(c *gin.Context) {
    books := []models.Book
    models.DB.Find(&books)

    c.JSON(http.StatusOK, gin.H{"data": books})
}
