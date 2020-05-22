package reader

import 	"github.com/gin-gonic/gin"


func GetInfos  (ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "实现普通的get请求",
	})
}