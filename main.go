// main.go
package main

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Configlist struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config Configlist

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = Configlist{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
