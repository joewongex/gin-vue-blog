package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
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
		log.Fatal("读取配置文件出错，请检查配置路径：", err)
	}
	loadServer(file)
	loadDatabase(file)
}

func loadServer(file *ini.File) {
	server := file.Section("server")
	AppMode = server.Key("AppMode").MustString("debug")
	HttpPort = server.Key("HttpPort").MustString(":3000")
}

func loadDatabase(file *ini.File) {
	database := file.Section("database")
	Db = database.Key("Db").MustString("mysql")
	DbHost = database.Key("DbHost").MustString("localhost")
	DbPort = database.Key("DbPort").MustString("3306")
	DbUser = database.Key("DbUser").MustString("ginvblog")
	DbPassword = database.Key("DbPassword").MustString("admin123")
	DbName = database.Key("DbName").MustString("ginvblog")
}
