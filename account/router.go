package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Router(m *gin.Engine) {
	r := m.Group("/account")
	r.GET("/", status)
}
