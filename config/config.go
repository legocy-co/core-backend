package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"legocy-go/pkg/helpers"
	"os"
	"strconv"
)

var appConfigInstance *AppConfig // private singleton variable

type AppConfig struct {
	DbConf    DatabaseConfig `yaml:"database" json:"database"`
	JwtConf   JWTConfig      `yaml:"jwt" json:"jwt"`
	KafkaConf KafkaConfig    `yaml:"kafka" json:"kafka"`

	S3Port string `json:"s3_port"`

	EventNotifierPort   string `json:"event_notifier_port"`
	EventNotifierChatID int    `json:"event_notifier_chat_id"`
}

func GetAppConfig() *AppConfig {
	return appConfigInstance
}

func SetAppConfig(cfg *AppConfig) error {
	if appConfigInstance != nil {
		return ErrConfigAlreadyExists
	}

	appConfigInstance = cfg
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

	return SetAppConfig(&cfg)
}

func SetupFromEnv() error {
	dbHost := os.Getenv("DB_HOST")
	logrus.Printf("DB_HOST = %v", dbHost)
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	logrus.Printf("DB_PORT = %v", dbPort)
	dbUser := os.Getenv("DB_USER")
	logrus.Printf("DB_USER = %v", dbUser)
	dbPassword := os.Getenv("DB_PASSWORD")
	logrus.Printf("DB_PASSWORD = %v", dbPassword)
	dbDatabaseName := os.Getenv("DB_DATABASE_NAME")
	loadFixtures := os.Getenv("DB_LOAD_FIXTURES") == "true"

	dbConfig := DatabaseConfig{
		Hostname:     dbHost,
		Port:         dbPort,
		DbName:       dbDatabaseName,
		DbUser:       dbUser,
		DbPassword:   dbPassword,
		LoadFixtures: loadFixtures,
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtAccessTokenLifetimeHours, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFETIME"))

	jwtConfig := JWTConfig{
		SecretKey:          jwtSecretKey,
		AccesTokenLifeTime: jwtAccessTokenLifetimeHours,
	}

	kafkaUri := os.Getenv("KAFKA_URI")

	kafkaConfig := KafkaConfig{kafkaUri}

	s3Port := os.Getenv("S3_PORT")

	eventNotifierPort := "<REMOVE ME>"
	eventNotifierChatID := 0

	appConfig := AppConfig{
		DbConf:              dbConfig,
		JwtConf:             jwtConfig,
		KafkaConf:           kafkaConfig,
		S3Port:              s3Port,
		EventNotifierPort:   eventNotifierPort,
		EventNotifierChatID: eventNotifierChatID,
	}

	return SetAppConfig(&appConfig)
}