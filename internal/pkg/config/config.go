package config

var appConfigInstance *AppConfig // private singleton variable

type AppConfig struct {
	BaseURL string `json:"base_url"`

	DbConf    DatabaseConfig `yaml:"database" json:"database"`
	JwtConf   JWTConfig      `yaml:"jwt" json:"jwt"`
	KafkaConf KafkaConfig    `yaml:"kafka" json:"kafka"`

	S3Host     string `json:"s3_host"`
	S3Port     string `json:"s3_port"`
	CDNBaseURL string `json:"cdn_base_url"`

	GoogleClientID string `json:"google_client_id"`

	FacebookAppID         string `json:"facebook_app_id"`
	FacebookSecret        string `json:"facebook_secret"`
	FacebookCallbackURL   string `json:"facebook_callback_url"`
	FacebookSessionSecret string
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
	URI             string `json:"uri"`
	ConsumerGroupId string `json:"consumer_group_id"`
}

func GetDBConfig() *DatabaseConfig {
	cfg := GetAppConfig()
	if cfg == nil {
		return nil
	}

	return &cfg.DbConf
}
