package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configs *Config

type Config struct {
	Version      string
	ServiceName  string
	HTTPPort     int
	JWTSecretKey string
	DB           *DBConfig
}
type DBConfig struct {
	Host          string `mapstructure:"DB_HOST" json:"host"`
	Port          int    `mapstructure:"DB_PORT" json:"port"`
	Name          string `mapstructure:"DB_NAME" json:"name"`
	User          string `mapstructure:"DB_USER" json:"user"`
	Password      string `mapstructure:"DB_PASSWORD" json:"password"`
	EnableSSLMode bool   `mapstructure:"DB_SSL_MODE" json:"enable_ssl_mode"`
}

func loanConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load the env variavles", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("version is required")
	}

	ServiceName := os.Getenv("SERVICE_NAME")
	if ServiceName == "" {
		log.Fatal("Service Name is required")
	}

	HTTPPort := os.Getenv("HTTP_PORT")
	if HTTPPort == "" {
		log.Fatal("HTTP Port is required")
	}

	port, err := strconv.ParseInt(HTTPPort, 10, 64)
	if err != nil {
		log.Fatal("HTTP_PORT must be a number")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is required")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST is required")
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		log.Fatal("DB_PORT is required")
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal("DB_PORT must be a number")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER is required")
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	sslMode := os.Getenv("DB_SSL_MODE")

	enableSSL := false
	if sslMode == "enable" || sslMode == "require" {
		enableSSL = true
	}
	configs = &Config{
		Version:      version,
		ServiceName:  ServiceName,
		HTTPPort:     int(port),
		JWTSecretKey: jwtSecretKey,
		DB: &DBConfig{
			Host:          dbHost,
			Port:          dbPort,
			Name:          dbName,
			User:          dbUser,
			Password:      dbPassword,
			EnableSSLMode: enableSSL,
		},
	}
}

func GetConfig() *Config {

	if configs == nil {
		loanConfig()
	}

	return configs
}
