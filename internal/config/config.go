package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"legocy-go/pkg/helpers"
)

var appConf *AppConfig // private singleton variable
var DefaultJWTConfig *JWTConfig = &JWTConfig{SecretKey: "test12345", AccesTokenLifeTime: 3}

type DatabaseConfig struct {
	Hostname   string `yaml:"hostname" json:"hostname"`
	Port       int    `yaml:"port" json:"port"`
	DbName     string `yaml:"db_name" json:"db_name"`
	DbUser     string `yaml:"db_user" json:"db_user"`
	DbPassword string `yaml:"db_password" json:"db_password"`
}

type JWTConfig struct {
	SecretKey          string `yaml:"secret_key" json:"secret_key"`
	AccesTokenLifeTime int    `yaml:"acces_tokern_lifetime_hours" json:"acces_token_lifetime_hours"`
}

type AppConfig struct {
	DbConf  DatabaseConfig `yaml:"database" json:"database"`
	JwtConf JWTConfig      `yaml:"jwt" json:"jwt"`
}

func GetAppConfig() *AppConfig {
	return appConf
}

func GetDBConfig() *DatabaseConfig {
	cfg := GetAppConfig()
	if cfg == nil {
		return nil
	}

	return &cfg.DbConf
}

func GetJWTConfig() *JWTConfig {
	cfg := GetAppConfig()
	if cfg == nil {
		return DefaultJWTConfig
	}

	return &cfg.JwtConf
}

func SetAppConfig(cfg *AppConfig) error {
	if appConf != nil {
		return ErrConfigAlreadyExists
	}

	appConf = cfg
	return nil
}

func SetupFromJSON(fp string) error {
	var cfg AppConfig

	if fileExists := helpers.FileExists(fp); !fileExists {
		return ErrConfigFileDoesNotExist
	}

	raw, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	fmt.Println(raw)

	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return err
	}

	SetAppConfig(&cfg)
	return nil
}
