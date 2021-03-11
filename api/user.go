package api

import (
	"cstore/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func Register(c *gin.Context) {

	m := make(map[string]string)
	m["cly"] = "hero"
	user := model.User{
		Model: gorm.Model{

			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserName:       "cly",
		Email:          "571821151@qq.com",
		PasswordDigest: "123123",
		Nickname:       "cly",
		Status:         "1",
		Limit:          0,
		Avatar:         "1",
	}
	if err := model.DB.Create(user).Error; err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, m)

}
