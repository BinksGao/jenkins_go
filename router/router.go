package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"gojenkis/reader"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	rUser := router.Group("/v1")
	{
		rUser.GET("/getInfo", reader.GetInfos)     // 获取banners
	}
	return router
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
			c.Set("content-type", "text/html")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}