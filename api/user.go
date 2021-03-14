package api

import (
	"cstore/common"
	svs "cstore/svs"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var userlimit svs.UserLitmit
	if err := c.ShouldBind(&userlimit); err != nil {

		c.JSON(500, err)
		return
	}
	rep := userlimit.List()
	c.JSON(200, rep)

}

func Login(c *gin.Context) {
	var login svs.LoginForm
	if err := c.ShouldBind(&login); err != nil {

		c.JSON(500, err)
		return
	}
	rep := login.Login()
	c.JSON(200, rep)

}

/*
创建新用户
*/
func CreateUser(c *gin.Context) {
	//session := sessions.Default(c)
	//userId := session.Get("userId")
	var service svs.UserSvs
	if err := c.ShouldBind(&service); err == nil {
		c.JSON(200, service.Register("1", nil))
	} else {
		c.JSON(200, common.ErrorResponse(err))
	}
}
