package main

type MCPConfig struct {
	Server *ServerStruct
}

type ServerStruct struct {
	AppName    string
	AppVersion string
}

func New() *MCPConfig {
	return &MCPConfig{
		Server: &ServerStruct{
			AppName:    GetEnvString("APP_NAME", "weather"),
			AppVersion: GetEnvString("APP_VERSION", "1.0"),
		},
	}
}
