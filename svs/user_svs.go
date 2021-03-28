package svs

import (
	"cstore/cache"
	"cstore/common"
	"cstore/model"
	"cstore/pkg/e"
	"cstore/pkg/logging"
	"cstore/serializer"
	"encoding/json"
	"time"
)

type UserSvs struct {
	Nickname  string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Challenge string `form:"challenge" json:"challenge"`
	Validate  string `form:"validate" json:"validate"`
	Seccode   string `form:"seccode" json:"seccode"`
}

type UserLitmit struct {
	Limit uint32 `form:"limit" json:"limit" `
	Start uint32 `form:"start" json:"start"`
}

/*用户登录form*/
type LoginForm struct {
	UserName string `form:"username" json:"username" `
	Password string `form:"password" json:"password"`
}

func (l *LoginForm) Login(now time.Time) *serializer.Response {

	code := common.SUCCESS
	if l.UserName == "" || l.Password == "" {
		code = common.ERROR_NOT_EXIST_USER
		return &serializer.Response{
			Status: code,
			Msg:    common.GetMsg(code),
		}
	}
	user := model.User{
		UserName: l.UserName,
	}

	println("user.SetPassword(l.Password)", time.Since(now).Milliseconds())
	user.SetPassword(l.Password)
	println("time.Now()", time.Since(now).Milliseconds())
	get := cache.RedisClient.Get(user.UserName)
	println("get cache", time.Since(now).Milliseconds())
	if get.Val() == "" {
		logging.Info("no cache")

		var count int
		if err := model.DB.Model(&user).Where("user_name=? and password_digest=?", l.UserName, user.PasswordDigest).Count(&count).Error; err != nil {
			return common.ErrorResponse(err)
		}
		marshal, err := json.Marshal(user)
		if err != nil {
			logging.Error(err)
		}

		set := cache.RedisClient.Set(user.UserName, string(marshal), 12*time.Hour)
		println(set)
		token, err := common.GenerateToken(user.UserName, user.PasswordDigest, 0)

		if err != nil {
			logging.Error(err)
			return common.ErrorResponse(err)
		}
		m := map[string]string{"token": token}
		return &serializer.Response{
			Status: code,
			Data:   m,
		}
	}
	u := get.Val()

	err := json.Unmarshal([]byte(u), &user)
	println("json Unmarshal", time.Since(now).Milliseconds())

	if err != nil {
		logging.Error(err)

	}
	return &serializer.Response{
		Status: code,
		Data:   user,
	}
}

func (s *UserSvs) Vaild(id, status interface{}) *serializer.Response {
	var count int
	err := model.DB.Model(&model.User{}).Where("nickname = ?", s.Nickname).Count(&count).Error
	if err != nil {
		return &serializer.Response{
			Status: e.ERROR_DATABASE,
			Msg:    common.GetMsg(e.ERROR_DATABASE),
		}
	}
	return nil
}
func (s *UserSvs) Register(userID, status interface{}) *serializer.Response {
	user := model.User{
		UserName: s.UserName,
		Nickname: s.Nickname,
		Status:   model.Active,
	}
	code := common.SUCCESS
	if res := s.Vaild(userID, status); res != nil {
		return res
	}
	if err := user.SetPassword(s.Password); err != nil {
		logging.Error(err)
		code = common.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Error:  common.GetMsg(code),
		}
	}
	user.Avatar = "img/avatar/avatar1.jpg"
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Error(err)
		code = common.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    common.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return &serializer.Response{
		Status: code,

		Msg: e.GetMsg(code),
	}
}

func (l *UserLitmit) List() *serializer.Response {
	var users []model.User
	code := common.SUCCESS

	if err := model.DB.Model(model.User{}).Offset(l.Start).Limit(l.Limit).Find(&users).Error; err != nil {
		code = common.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    common.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   users,
		Msg:    common.GetMsg(common.SUCCESS),
	}

}
