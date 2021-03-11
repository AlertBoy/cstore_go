package server

import (
	"cstore/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//him
	gin.SetMode(gin.DebugMode)
	v1 := r.Group("/api/v1")
	{
		v1.GET("user/register", api.Register)
	}

	return r
}
