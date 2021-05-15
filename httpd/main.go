package main


import (
	"GoApi/httpd/handler"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.GET("/testEndpoint", handler.TestGet)

	r.Run()
}