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
		case 2002:
		case 2006:
			c.JSON(http.StatusNotFound, gin.H{"error": appErr.Message})
		case 2008:
			c.JSON(http.StatusBadRequest, gin.H{"error": appErr.Message})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": appErr.Message})
		}
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
