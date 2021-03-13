package svs

import (
	"cstore/common"
	"cstore/model"
	"cstore/pkg/e"
	"cstore/pkg/logging"
	"cstore/serializer"
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
