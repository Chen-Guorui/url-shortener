package config

import (
	"encoding/json"
	"os"
)

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type LogConfig struct {
	Path    string
	Server  string
	Request string
}

type C struct {
	MySQL MySQLConfig
	Port  int
	Log   LogConfig
}

var (
	Config C
)

func init() {
	data, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
}
