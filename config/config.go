package config

import (
	"os"

	"github.com/spf13/viper"
)

type CronConfig struct {
	Enabled bool
	Spec    string
}

type Config struct {
	Cron CronConfig
}

func NewConfig() (*Config, error) {
	c := &Config{}
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}
	// Populate Cron from viper
	c.Cron.Enabled = viper.GetBool("cron.enabled")
	c.Cron.Spec = viper.GetString("cron.spec")
	if err := c.loadConfig(); err != nil {
		return nil, err
	}
	return c, nil
}

func (*Config) loadConfig() error {
	return viper.ReadInConfig()
}

func (*Config) GetString(key string) string {
	return viper.GetString(key)
}

func (*Config) GetInt(key string) int {
	return viper.GetInt(key)
}

func (*Config) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (*Config) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (*Config) Getenv(key string) string {
	return os.Getenv(key)
}
