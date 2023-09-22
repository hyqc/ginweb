package config

import (
	"ginweb/pkg/logger"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type DB struct {
	Demo *gorm.DB
}

var (
	AppConfig = &Config{}
	AppDB     = &DB{}
	AppLogger = logger.NewLogger()
)

func ParseConfig(name string) error {
	body, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(body, AppConfig); err != nil {
		return err
	}
	return nil
}
