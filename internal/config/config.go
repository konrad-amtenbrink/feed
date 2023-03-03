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
	Username         string
	Password         string
	Host             string
	Port             int
	Database         string
	ConnectTimeout   int
	DisableSslMode   bool
	ConnectionString string
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
			Username:         getEnv("DB_USERNAME"),
			Password:         getEnv("DB_PASSWORD"),
			Host:             getEnv("DB_HOST"),
			Port:             parseInt(getEnv("DB_PORT"), false),
			Database:         getEnv("DB_DATABASE"),
			ConnectTimeout:   parseInt(getEnv("DB_CONNECT_TIMEOUT"), false),
			DisableSslMode:   parseBool(getEnv("DB_DISABLE_SSL_MODE"), false),
			ConnectionString: getEnv("DATABASE_URL"),
		},
		Logging: LoggingConfig{
			Level:               mustGetEnv("LOG_LEVEL"),
			EnableReportCaller:  parseBool(mustGetEnv("LOG_REPORT_CALLER"), true),
			EnableTextFormatter: parseBool(mustGetEnv("LOG_TEXT_FORMATTER"), true),
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

// Try to parse an environment variable. If it does not exist, return the default value.
func getEnv(env string) string {
	val, exists := os.LookupEnv(env)

	if !exists {
		return ""
	}

	return val
}

// Parse a boolean value, "true" or "false", or fatal the application.
func parseBool(val string, mustParse bool) bool {
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		if mustParse {
			log.WithError(err).Fatalf("failed to parse bool with value %s", val)
		} else {
			return false
		}
	}

	return boolVal
}

// Parse an integer value, or fatal the application.
func parseInt(val string, mustParse bool) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		if mustParse {
			log.WithError(err).Fatalf("failed to parse int with value %s", val)
		} else {
			return 0
		}
	}

	return intVal
}
