package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DB DbConfig
}

type DbConfig struct {
	Migration  Migration
	Connection Connection
}

type Migration struct {
	Folder           string
	ConnectionString string
	Enable           bool
}

type Connection struct {
	URL string
}

func GetConfig(configName string) *Config {
	project_path, _ := os.Getwd()
	viper.SetConfigName(configName)
	viper.AddConfigPath(project_path)
	viper.ReadInConfig()

	return &Config{
		DbConfig{
			Connection: Connection{
				URL: viper.GetString("db.connection.url"),
			},
			Migration: Migration{
				Folder:           viper.GetString("db.migrations.folder"),
				ConnectionString: viper.GetString("db.migrations.connectionString"),
				Enable:           viper.GetBool("db.migrations.enable"),
			},
		},
	}
}
