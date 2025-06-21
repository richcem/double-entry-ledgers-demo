package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	JWTSecret string `yaml:"jwt_secret"`
}

// Load 加载配置文件
func Load() (*Config, error) {
	// 默认配置
	config := &Config{
		JWTSecret: "your-secret-key-here", // 实际生产环境中应使用更安全的密钥
	}

	// 如果配置文件存在，则加载配置文件
	if _, err := os.Stat("config.yaml"); err == nil {
		data, err := os.ReadFile("config.yaml")
		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(data, config); err != nil {
			return nil, err
		}
	}

	return config, nil
}
