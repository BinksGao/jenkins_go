package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Get请求成功!")
	})
	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}