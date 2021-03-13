package server

import (
	"cstore/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//him
	gin.SetMode(gin.DebugMode)
	// 用户登录
	v1 := r.Group("/api/v1")
	{
		v1.POST("user", api.CreateUser)
		v1.GET("user", api.GetUser)

	}

	return r
}
