package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MQTTClientID string `mapstructure:"MQTT_CLIENT_ID"`
	MQTTBroker   string `mapstructure:"MQTT_BROKER"`
	KafkaBroker  string `mapstructure:"KAFKA_BROKER"`
	DatabaseUrl  string `mapstructure:"DATABASE_URL"`
	DeviceTopic  string `mapstructure:"DEVICE_TOPIC"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file couldn't be found, using default values or environment variables.")
		} else {
			return nil, fmt.Errorf("config file couldn't be read: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("config cannot be unmarshaled: %w", err)
	}

	return &config, nil
}
