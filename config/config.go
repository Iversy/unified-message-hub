package config

import (
	"fmt"
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Kafka    KafkaConfig    `yaml:"kafka"`
	Web      Web            `yaml:"web"`
	VK       VK             `yaml:"vk"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type KafkaConfig struct {
	Host                   string `yaml:"host"`
	Port                   string `yaml:"port"`
	MessageCreateTopicName string `yaml:"message_create_topic_name"`
}

type Web struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type VK struct {
	GroupID     int           `yaml:"group_id"`
	AdminID     int           `yaml:"admin_id"`
	WelcomeText string        `yaml:"welcome_text"`
	Timeout     time.Duration `yaml:"timeout"` // for updates (s)
	Delay       time.Duration `yaml:"delay"`   // for subsequent requests (ms)
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}
