package utils

import (
    "log"

    "github.com/gin-gonic/gin"
)

func TreatErrorResponse(context *gin.Context, message string) {
    log.Println(message)
    context.JSON(400, gin.H{"msg": message})
}
