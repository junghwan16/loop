package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	DBURL        string `mapstructure:"DB_URL"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`
}

func Load() *Config {
	viper.SetConfigName("config") // config 파일 이름 (e.g., config.yaml)
	viper.SetConfigType("env")    // 또는 "env"로 .env 사용 가능
	viper.AddConfigPath(".")      // 현재 디렉토리
	viper.AddConfigPath("./config")

	// 환경 변수 우선 적용 (e.g., SERVER_PORT 환경 변수가 있으면 오버라이드)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Config file not found, using env vars: %v\n", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("unable to decode config: %w", err))
	}

	// 필수 값 검증 (예: OpenAI 키 없으면 패닉)
	if cfg.OpenAIAPIKey == "" {
		panic("OPENAI_API_KEY is required")
	}
	if cfg.DBURL == "" {
		panic("DB_URL is required")
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080" // 기본값
	}

	return &cfg
}
