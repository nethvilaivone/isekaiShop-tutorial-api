package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Database *Database `mapstructure:"database" validate:"required"`
		Server   *Server   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2   `mapstructure:"oauth2" validate:"required"`
		State    *State    `mapstructure:"state" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"` // Changed from int to string
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

	Server struct {
		Port         int          `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
	}

	OAuth2 struct {
		PlayerRedirectUrl string   `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string   `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientId          string   `mapstructure:"clientId" validate:"required"`
		ClientSecret      string   `mapstructure:"clientSecret" validate:"required"`
		Endpoints         endpoint `mapstructure:"endpoints" validate:"required"`
		Scopes            []string `mapstructure:"scopes" validate:"required"`
		UserInfoUrl       string   `mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl         string   `mapstructure:"revokeUrl" validate:"required"`
	}

	endpoint struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}

	State struct {
		Secret    string        `mapstructure:"secret" validate:"required"`
		ExpiresAt time.Duration `mapstructure:"expiresAt" validate:"required"`
		Isuser    string        `mapstructure:"isuser" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		configInstance = &Config{} // Initialize the config instance

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".") // Also look in current directory
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(configInstance); err != nil {
			panic(err)
		}

		validating := validator.New()
		if err := validating.Struct(configInstance); err != nil {
			panic(err) // Changed from panic(validating)
		}
	})

	return configInstance
}
