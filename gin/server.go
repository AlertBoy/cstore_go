package server

import (
	"cstore/api"
	"cstore/middleware"
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
		v1.POST("user/login", api.Login)

	}
	v2 := r.Group("/api/v12")
	{
		v2.POST("mq", api.MqTest)
		v2.GET("mq", api.MqConsume)
	}
	authed := r.Group("/").Use(middleware.Cors(), middleware.JWT())
	{
		authed.POST("user", api.GetUser)
	}

	return r
}
