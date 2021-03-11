package config

import (
	"cstore/model"
)

func Init() {
	//godotenv.Load()
	//// 读取翻译文件
	//if err := loadLocal("conf/locales/zh-cn.yaml"); err != nil {
	//	logging.Info(err)
	//	panic(err)
	//}
	// 连接数据库
	model.Database("root:123456@/cstore?charset=utf8&parseTime=True&loc=Local")
	//cache.Redis()
}
