package config

import (
	"strings"
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort int

	HttpRateLimitRequest int
	HttpRateLimitTime    time.Duration

	JwtSecretKey string
	JwtTTL       time.Duration

	PaginationLimit int

	MysqlUser            string
	MysqlPassword        string
	MysqlHost            string
	MysqlPort            int
	MysqlDatabase        string
	MysqlMaxIdleConns    int
	MysqlMaxOpenConns    int
	MysqlConnMaxLifetime time.Duration

	RedisPassword string
	RedisHost     string
	RedisPort     int
	RedisDatabase int
	RedisPoolSize int
	RedisTTL      time.Duration
}

func load() Config {
	vp := viper.New()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.AutomaticEnv()

	return Config{
		AppPort:              vp.GetInt("app.port"),
		HttpRateLimitRequest: vp.GetInt("http.rate.limit.request"),
		HttpRateLimitTime:    vp.GetDuration("http.rate.limit.time"),
		JwtSecretKey:         vp.GetString("jwt.secret.key"),
		JwtTTL:               vp.GetDuration("jwt.ttl"),
		PaginationLimit:      vp.GetInt("pagination.limit"),
		MysqlUser:            vp.GetString("mysql.user"),
		MysqlPassword:        vp.GetString("mysql.password"),
		MysqlHost:            vp.GetString("mysql.host"),
		MysqlPort:            vp.GetInt("mysql.port"),
		MysqlDatabase:        vp.GetString("mysql.database"),
		MysqlMaxIdleConns:    vp.GetInt("mysql.max.idle.conns"),
		MysqlMaxOpenConns:    vp.GetInt("mysql.max.open.conns"),
		MysqlConnMaxLifetime: vp.GetDuration("mysql.conn.max.lifetime"),
		RedisPassword:        vp.GetString("redis.password"),
		RedisHost:            vp.GetString("redis.host"),
		RedisPort:            vp.GetInt("redis.port"),
		RedisDatabase:        vp.GetInt("redis.database"),
		RedisPoolSize:        vp.GetInt("redis.pool.size"),
		RedisTTL:             vp.GetDuration("redis.ttl"),
	}
}

var config = load()

func Cfg() *Config { return &config }
