package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configName string = "config"
	configPath string = "./config/.env" // config.yaml
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"pgdb"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		HttpAddr         string `env-required:"true" yaml:"addr"          env:"HTTP_ADDR"`
		HttpReadTimeout  int    `env-required:"true" yaml:"read_timeout"  env:"HTTP_READTIMEOUT"`
		HttpWriteTimeout int    `env-required:"true" yaml:"write_timeout" env:"HTTP_WRITETIMEOUT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		Host        string `env-required:"true" yaml:"host"         env:"PG_HOST"`
		Username    string `env-required:"true" yaml:"username"     env:"PG_USERNAME"`
		Password    string `env-required:"true" yaml:"password"     env:"PG_PASSWORD"`
		Port        int    `env-required:"true" yaml:"port"         env:"PG_PORT"`
		DBName      string `env-required:"true" yaml:"dbname"       env:"PG_DBNAME"`
		ConnTimeout int    `env-required:"true" yaml:"conn_timeout" env:"PG_CONN_TIMEOUT"`
		PoolMax     int    `env-required:"true" yaml:"pool_max"     env:"PG_POOL_MAX"`
		//URL         string `env-required:"true"                     env:"PG_URL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
