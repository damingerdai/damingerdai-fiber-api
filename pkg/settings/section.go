package settings

var sections = make(map[string]any)

type ServerSettingS struct {
	RunMode  string
	HttpPort string
}


type AppSettingS struct {
	Server ServerSettingS
}


