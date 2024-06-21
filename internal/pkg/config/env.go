package config

import (
	"os"
	"strconv"
)

func SetupFromEnv() error {

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "localhost"
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
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
	jwtRefreshTokenLifetimeHours, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_LIFETIME"))

	jwtConfig := JWTConfig{
		SecretKey:            jwtSecretKey,
		AccessTokenLifeTime:  jwtAccessTokenLifetimeHours,
		RefreshTokenLifeTime: jwtRefreshTokenLifetimeHours,
	}

	kafkaUri := os.Getenv("KAFKA_URI")
	kafkaConsumerGroupId := os.Getenv("KAFKA_CONSUMER_GROUP_ID")

	kafkaConfig := KafkaConfig{kafkaUri, kafkaConsumerGroupId}

	s3Host := os.Getenv("S3_HOST")
	s3Port := os.Getenv("S3_PORT")
	cdnBaseUrl := os.Getenv("CDN_BASE_URL")

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	fbSalt := os.Getenv("FACEBOOK_SECRET_KEY_SALT")

	appConfig := AppConfig{
		BaseURL:               baseUrl,
		DbConf:                dbConfig,
		JwtConf:               jwtConfig,
		KafkaConf:             kafkaConfig,
		S3Host:                s3Host,
		S3Port:                s3Port,
		CDNBaseURL:            cdnBaseUrl,
		GoogleClientID:        googleClientID,
		FacebookSecretKeySalt: fbSalt,
	}

	return SetAppConfig(&appConfig)
}
