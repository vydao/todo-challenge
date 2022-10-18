package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:",squash"`
	Token    Token    `mapstructure:",squash"`
}

// Database holds the functional configuration settings related to Datatase connection
type Database struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	MigrationPath string `mapstructure:"MIGRATION_PATH"`
}

// Token holds the functional configuration settings related to JWT token
type Token struct {
	SecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

// LoadConfig reads config from file or environment variables
func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	conf := Config{}
	if err := viper.ReadInConfig(); err != nil {
		return conf, err
	}
	err := viper.Unmarshal(&conf)
	return conf, err
}
