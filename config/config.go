package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DB      DBConfig
	Logging LoggingConfig
	AWS     AWSConfig
}

type LoggingConfig struct {
	Level               string
	EnableReportCaller  bool
	EnableTextFormatter bool
}

type DBConfig struct {
	Username       string
	Password       string
	Host           string
	Port           int
	Database       string
	ConnectTimeout int
	DisableSslMode bool
}

type AWSConfig struct {
	Region          string
	AccessKey       string
	SecretAccessKey string
	BucketName      string
}

type Auth struct {
	SecretKey string
}

// MustParseConfig loads the configuration from the environment
func MustParseConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithError(err).Warn("Could not load .env file")
	}

	config := Config{
		DB: DBConfig{
			Username:       mustGetEnv("DB_USERNAME"),
			Password:       mustGetEnv("DB_PASSWORD"),
			Host:           mustGetEnv("DB_HOST"),
			Port:           mustParseInt(mustGetEnv("DB_PORT")),
			Database:       mustGetEnv("DB_DATABASE"),
			ConnectTimeout: mustParseInt(mustGetEnv("DB_CONNECT_TIMEOUT")),
			DisableSslMode: mustParseBool(mustGetEnv("DB_DISABLE_SSL_MODE")),
		},
		Logging: LoggingConfig{
			Level:               mustGetEnv("LOG_LEVEL"),
			EnableReportCaller:  mustParseBool(mustGetEnv("LOG_REPORT_CALLER")),
			EnableTextFormatter: mustParseBool(mustGetEnv("LOG_TEXT_FORMATTER")),
		},
		AWS: AWSConfig{
			Region:          mustGetEnv("AWS_REGION"),
			AccessKey:       mustGetEnv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: mustGetEnv("AWS_SECRET_ACCESS_KEY"),
			BucketName:      mustGetEnv("AWS_BUCKET_NAME"),
		},
	}

	return config
}

// Parse an environment variable, or fatal the application.
func mustGetEnv(env string) string {
	val, exists := os.LookupEnv(env)

	if !exists {
		log.WithField("key", env).Fatalf("undefined environment variable")
	}

	return val
}

// Parse a boolean value, "true" or "false", or fatal the application.
func mustParseBool(val string) bool {
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		log.WithError(err).Fatalf("failed to parse bool with value %s", val)
	}

	return boolVal
}

// Parse an integer value, or fatal the application.
func mustParseInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.WithError(err).Fatalf("failed to parse int with value %s", val)
	}

	return intVal
}
