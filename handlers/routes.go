package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRoute(port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test route working!!", "port": port})
	}

}
