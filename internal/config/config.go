package config

import (
	"encoding/json"
	"io/ioutil"
	"legocy-go/pkg/helpers"
)

var AppConfigInstance *AppConfig // private singleton variable

type AppConfig struct {
	DbConf    DatabaseConfig `yaml:"database" json:"database"`
	JwtConf   JWTConfig      `yaml:"jwt" json:"jwt"`
	KafkaConf KafkaConfig    `yaml:"kafka" json:"kafka"`

	S3Port string `json:"s3_port"`

	EventNotifierPort   string `json:"event_notifier_port"`
	EventNotifierChatID int    `json:"event_notifier_chat_id"`
}

func GetAppConfig() *AppConfig {
	return AppConfigInstance
}

func SetAppConfig(cfg *AppConfig) error {
	if AppConfigInstance != nil {
		return ErrConfigAlreadyExists
	}

	AppConfigInstance = cfg
	return nil
}

type DatabaseConfig struct {
	Hostname     string `yaml:"hostname" json:"hostname"`
	Port         int    `yaml:"port" json:"port"`
	DbName       string `yaml:"db_name" json:"db_name"`
	DbUser       string `yaml:"db_user" json:"db_user"`
	DbPassword   string `yaml:"db_password" json:"db_password"`
	LoadFixtures bool   `yaml:"load_fixtures" json:"load_fixtures"`
}

type JWTConfig struct {
	SecretKey          string `yaml:"secret_key" json:"secret_key"`
	AccesTokenLifeTime int    `yaml:"acces_tokern_lifetime_hours" json:"acces_token_lifetime_hours"`
}

type KafkaConfig struct {
	URI string `json:"uri"`
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

func SetupFromJSON(fp string) error {
	var cfg AppConfig

	if fileExists := helpers.FileExists(fp); !fileExists {
		return ErrConfigFileDoesNotExist
	}

	raw, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return err
	}

	SetAppConfig(&cfg)
	return nil
}
