package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gostack-labs/adminx/bootstrap"
)

func init() {
	bootstrap.Initialize()
}

func main() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.NoRoute(func(ctx *gin.Context) {
		acceptStr := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 URL 和请求方法是否正确。"})
		}
	})

	r.Run(":9999")
}
