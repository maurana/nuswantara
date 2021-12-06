package config

import (
	"strings"
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort 				int
	HttpRateLimitRequest	int
	HttpRateLimitTime		time.Duration
	JwtSecretKey			string
	JwtTTL       t			time.Duration
}

func load() Config {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigFile(".env")
	err := vp.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file %s", err)
	}
	return Config{
		AppPort:              vp.GetInt("APP_PORT"),
		HttpRateLimitRequest: vp.GetInt("HTTP_RATE_LIMIT_REQUEST"),
		HttpRateLimitTime:    vp.GetDuration("HTTP_RATE_LIMIT_TIME"),
		JwtSecretKey:         vp.GetString("JWT_SCRET_KEY"),
		JwtTTL:               vp.GetDuration("JWT_TTL"),
	}
}

var config = load()
func Cfg() *Config { return &config }