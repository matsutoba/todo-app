package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		switch appErr.Code {
		case 2001:
			c.JSON(http.StatusConflict, gin.H{"error": appErr.Message})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
		}
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
