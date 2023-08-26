package main

import (
	"log"

	"github.com/dammingerdai/damingerdai-fiber-api/internal/routers"
	"github.com/dammingerdai/damingerdai-fiber-api/pkg/settings"
)

func init() {
	var err error

	appSetting, err := setupSetting()
	if err != nil {
		panic(err.Error())
	}
	log.Println(appSetting)
}

func main() {
	app := routers.NewRouter()

	app.Listen(":3000")

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
