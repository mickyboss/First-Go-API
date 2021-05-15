package handler


import (
	"net/http"
	"github.com/gin-gonic/gin"
)



func TestGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello": "im testing",
	})
}