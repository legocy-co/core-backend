package config

import (
	"encoding/json"
	"io/ioutil"
	"legocy-go/helpers"
)

type DatabaseConfig struct {
	Hostname   string `yaml:"hostname" json:"hostname"`
	Port       int    `yaml:"port" json:"port"`
	DbName     string `yaml:"db_name" json:"db_name"`
	DbUser     string `yaml:"db_user" json:"db_user"`
	DbPassword string `yaml:"db_password" json:"db_password"`
}

type JWTConfig struct {
	SecretKey string `yaml:"secret_key" json:"secret_key"`
}

type AppConfig struct {
	DbConf  DatabaseConfig `yaml:"database" json:"database"`
	JwtConf JWTConfig      `yaml:"jwt" json:"jwt"`
}

var appConf *AppConfig // private singleton variable
func GetAppConfig() *AppConfig {
	return appConf
}

func SetAppConfig(cfg *AppConfig) error {
	if appConf != nil {
		return ErrConfigAlreadyExists
	}

	appConf = cfg
	return nil
}

func SetupFromJSON(fp string) error {
	var cfg *AppConfig

	if fileExists := helpers.FileExists(fp); !fileExists {
		return ErrConfigFileDoesNotExist
	}

	raw, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, cfg)
	if err != nil {
		return err
	}

	appConf = cfg
	return nil
}
