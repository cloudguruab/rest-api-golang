package main

import (
    "net/http"
    
    "github.com/cloudguruab/rest-api-golang/controllers"
    "github.com/cloudguruab/rest-api-golang/models"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    models.ConnectDatabase()
    
    r.GET("/books", controllers.FindBooks)

    r.Run()
}
