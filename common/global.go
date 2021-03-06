package common

import (
	"cstore/serializer"
	"encoding/json"
	"fmt"
	validator "gopkg.in/go-playground/validator.v8"
)

const (
	SUCCESS                 = 200
	UPDATE_PASSWORD_SUCCESS = 201
	NOT_EXIST_IDENTIFIER    = 202
	ERROR                   = 500
	INVALID_PARAMS          = 400
	INVALID_TOKEN           = 401

	ERROR_EXIST_NICK           = 10001
	ERROR_EXIST_USER           = 10002
	ERROR_NOT_EXIST_USER       = 10003
	ERROR_NOT_COMPARE          = 10004
	ERROR_NOT_COMPARE_PASSWORD = 10005
	ERROR_FAIL_ENCRYPTION      = 10006
	ERROR_NOT_EXIST_PRODUCT    = 10007
	ERROR_NOT_EXIST_ADDRESS    = 10008
	ERROR_EXIST_FAVORITE       = 10009

	ERROR_AUTH_CHECK_TOKEN_FAIL       = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT    = 20002
	ERROR_AUTH_TOKEN                  = 20003
	ERROR_AUTH                        = 20004
	ERROR_AUTH_INSUFFICIENT_AUTHORITY = 20005
	ERROR_READ_FILE                   = 20006
	ERROR_SEND_EMAIL                  = 20007
	ERROR_CALL_API                    = 20008
	ERROR_UNMARSHAL_JSON              = 20009

	ERROR_DATABASE = 30001

	ERROR_OSS = 40001
	/*
	 应该放到配置文件中 viper is coming soon
	*/

)

var JWTSECRET = []byte("clyTest")
var msg = map[int]string{
	SUCCESS:                    "ok",
	UPDATE_PASSWORD_SUCCESS:    "修改密码成功",
	NOT_EXIST_IDENTIFIER:       "该第三方账号未绑定",
	ERROR:                      "fail",
	INVALID_PARAMS:             "请求参数错误",
	ERROR_EXIST_NICK:           "已存在该昵称",
	ERROR_EXIST_USER:           "已存在该用户名",
	ERROR_NOT_EXIST_USER:       "该用户不存在",
	ERROR_NOT_COMPARE:          "帐号密码错误",
	ERROR_NOT_COMPARE_PASSWORD: "两次密码输入不一致",
	ERROR_FAIL_ENCRYPTION:      "加密失败",
	ERROR_NOT_EXIST_PRODUCT:    "该商品不存在",
	ERROR_NOT_EXIST_ADDRESS:    "该收货地址不存在",
	ERROR_EXIST_FAVORITE:       "已收藏该商品",

	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token已超时",
	ERROR_AUTH_TOKEN:                  "Token生成失败",
	ERROR_AUTH:                        "Token错误",
	ERROR_AUTH_INSUFFICIENT_AUTHORITY: "权限不足",
	ERROR_READ_FILE:                   "读文件失败",
	ERROR_SEND_EMAIL:                  "发送邮件失败",
	ERROR_CALL_API:                    "调用接口失败",
	ERROR_UNMARSHAL_JSON:              "解码JSON失败",

	ERROR_DATABASE: "数据库操作出错，请重试",

	ERROR_OSS:     "OSS配置错误",
	INVALID_TOKEN: "token无效",
}

func GetMsg(code int) string {
	s, ok := msg[code]
	if !ok {
		return ""
	}
	return s
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) *serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := fmt.Sprintf("Field.%s", e.Field)
			tag := fmt.Sprintf("Tag.Valid.%s", e.Tag)
			return &serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return &serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return &serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
