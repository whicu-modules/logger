package config

import "log/slog"

type LumberjackConfig struct {
	Level    string `yaml:"level" env:"LOG_FILE_LEVEL" env-default:"info"`
	Path     string `yaml:"path" env:"LOG_FILE_PATH" env-default:""`
	Size     int    `yaml:"size" env:"LOG_FILE_SIZE" env-default:"128"`
	Compress bool   `yaml:"compress" env:"LOG_FILE_COMPRESS" env-default:"true"`
}

type Config struct {
	Level     string       `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
	AddSource bool         `yaml:"add_source" env:"LOG_ADD_SOURCE" env-default:"false"`
	Handler   slog.Handler `yaml:"-"`
	LumberjackConfig
}

func (c *Config) SetHandler(handler slog.Handler) {
	c.Handler = handler
}
