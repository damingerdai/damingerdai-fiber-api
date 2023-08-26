package main

import (
	"github.com/dammingerdai/damingerdai-fiber-api/global"
	"github.com/dammingerdai/damingerdai-fiber-api/internal/server"
	"github.com/dammingerdai/damingerdai-fiber-api/pkg/settings"
)

func init() {
	var err error

	appSetting, err := setupSetting()
	if err != nil {
		panic(err.Error())
	}
	global.AppSetting = appSetting
}

func main() {
	app := server.NewServer(global.AppSetting.Server)

	app.Run()

}

func setupSetting() (*settings.AppSettingS, error) {
	st, err := settings.NewSetting()
	if err != nil {
		return nil, err
	}
	var appSetting settings.AppSettingS
	st.ReadAllSection(&appSetting)
	return &appSetting, nil
}
