package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	configPathFromRoot = "/config"
)

type RootConfig struct {
	APIServerAddress string `yaml:"apiServerAddress"`
	TokenPath        string `yaml:"tokenPath"`
}

var Config RootConfig

func InitConfig() {
	configPath := os.Getenv("ENV_MOAICTL_ROOT") + configPathFromRoot
	viperConfig := viper.New()
	viperConfig.AddConfigPath(configPath)
	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("yaml")

	if err := viperConfig.ReadInConfig(); err != nil {
		fmt.Println("Error reading viper config", err)
		panic("failed to read viper config")
	}

	viperConfig.SetEnvPrefix("ENV_MOAI")

	replacer := strings.NewReplacer(".", "_")
	viperConfig.SetEnvKeyReplacer(replacer)

	viperConfig.AutomaticEnv()

	if err := viperConfig.Unmarshal(&Config); err != nil {
		fmt.Println("Error unmarshalling viper config", err)
		panic("failed to unmarshal viper config")
	}
}
