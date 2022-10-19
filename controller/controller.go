package controller

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func SampleFunc(c *gin.Context) {
    log.Println("I am sampleFunc")
    c.JSON(http.StatusOK, gin.H{"message": "konnnitiha"})
}