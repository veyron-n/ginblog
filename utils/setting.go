package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Print("配置文件读取错误,请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("server").Key("Db").MustString("mysql")
	DbHost = file.Section("server").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("server").Key("DbPort").MustString("3307")
	DbUser = file.Section("server").Key("DbUser").MustString("root")
	DbPassword = file.Section("server").Key("DbPassword").MustString("MYSQL")
	DbName = file.Section("server").Key("DbName").MustString("ginblog")
}
