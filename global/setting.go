package global

import (
	"service-template/internal/model"
	"service-template/pkg/settings"
	"time"
)

var (
	ServerSetting *settings.ServerSetting
	DatabaseSetting *settings.DatabaseSetting
	AppSetting *settings.AppSetting
	JWTSetting *settings.JWTSetting
)

// DBConnectSetting 配置数据库连接
func DBConnectSetting() error {
	var err error
	DBEngine, err = model.NewDBEngine(DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// SetupSetting 配置热更新
func SetupSetting() error {
	// 服务器基本配置
	setting, err := settings.NewSetting("configs/")
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	ServerSetting.WriteTimeout *= time.Second
	ServerSetting.ReadTimeout *= time.Second

	// 数据库配置
	err = setting.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	// App配置
	err = setting.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}

	// 配置JWT
	err = setting.ReadSection("JWT", &JWTSetting)
	JWTSetting.Expire *= time.Second

	return nil
}
