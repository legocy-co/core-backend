package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var appConfigInstance *AppConfig // private singleton variable

type AppConfig struct {
	BaseURL string `json:"base_url"`

	DbConf    DatabaseConfig `yaml:"database" json:"database"`
	JwtConf   JWTConfig      `yaml:"jwt" json:"jwt"`
	KafkaConf KafkaConfig    `yaml:"kafka" json:"kafka"`

	S3Host string `json:"s3_host"`
	S3Port string `json:"s3_port"`
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
	SecretKey            string `yaml:"secret_key" json:"secret_key"`
	AccessTokenLifeTime  int    `yaml:"access_token_lifetime_hours" json:"access_token_lifetime_hours"`
	RefreshTokenLifeTime int    `json:"refresh_token_life_time_hours"`
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

func SetupFromEnv() error {

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "localhost"
	}

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
	jwtRefreshTokenLifetimeHours, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_LIFETIME_HOURS"))

	jwtConfig := JWTConfig{
		SecretKey:            jwtSecretKey,
		AccessTokenLifeTime:  jwtAccessTokenLifetimeHours,
		RefreshTokenLifeTime: jwtRefreshTokenLifetimeHours,
	}

	kafkaUri := os.Getenv("KAFKA_URI")
	logrus.Printf("KAFKA_URL = %v", kafkaUri)

	kafkaConfig := KafkaConfig{kafkaUri}

	s3Host := os.Getenv("S3_HOST")
	s3Port := os.Getenv("S3_PORT")

	appConfig := AppConfig{
		BaseURL:   baseUrl,
		DbConf:    dbConfig,
		JwtConf:   jwtConfig,
		KafkaConf: kafkaConfig,
		S3Host:    s3Host,
		S3Port:    s3Port,
	}

	return SetAppConfig(&appConfig)
}
