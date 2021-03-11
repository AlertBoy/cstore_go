//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 15:01:44
 * @LastEditors: congz
 * @LastEditTime: 2020-07-19 14:58:56
 */
package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Admin 管理员模型
type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

// SetPassword 设置密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 封面地址
func (admin *Admin) AvatarURL() string {

	return ""
}
