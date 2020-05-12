package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
)

var (
	RunMode  string
	HttpPort int
)

type Config struct {
	RunMode string `toml:"RUN_MODE"`
	App     struct {
		PageSize  int    `toml:"PAGE_SIZE"`
		JwtSecret string `toml:"JWT_SECRET"`
	} `toml:"app"`
	Server struct {
		HttpPort     int `toml:"HTTP_PORT"`
		ReadTimeOut  int `toml:"READ_TIMEOUT"`
		WriteTimeOut int `toml:"WRITE_TIMEOUT"`
	} `toml:"server"`
	Database struct {
		Type     string `toml:"TYPE"`
		User     string `toml:"USER"`
		Password string `toml:"PASSWORD"`
		Host     string `toml:"HOST"`
		Port     string `toml:"PORT"`
		DbName   string `toml:"DB_NAME"`
	} `toml:"database"`
}

func GetConfig() Config {
	var cg Config
	filePath, _ := filepath.Abs("./config/conf.toml")
	if _, err := toml.DecodeFile(filePath, &cg); err != nil {
		log.Fatal(err)
	}
	return cg
}

func init() {
	var cg Config
	filePath, _ := filepath.Abs("./config/conf.toml")
	if _, err := toml.DecodeFile(filePath, &cg); err != nil {
		log.Fatal(err)
	}
	RunMode = cg.RunMode
	HttpPort = cg.Server.HttpPort
}
