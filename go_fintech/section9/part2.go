package section9

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLdriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("./section9/config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLdriver: cfg.Section("db").Key("driver").String(),
	}
}

func Main2() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLdriver, Config.SQLdriver)
}
