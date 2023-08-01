package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DEBUG                 bool
	PAPERLESS_AUTH_TOKEN  string
	APP_PORT              int
	PAPERLESS_EXPENSE_TAG string
	PAPERLESS_INCOME_TAG  string
	PAPERLESS_URL         string
	PAPERLESS_UNSAFE_SSL  bool
	FRONTEND_URL          string
	REDIS_USER            string
	REDIS_PASSWORD        string
	REDIS_DB              string
	REDIS_ADDRESS         string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New() (*Config, error) {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("No .env file found.\n")
	}

	config := Config{
		PAPERLESS_EXPENSE_TAG: getEnv("PAPERLESS_EXPENSE_TAG", "expense"),
		PAPERLESS_INCOME_TAG:  getEnv("PAPERLESS_INCOME_TAG", "income"),
	}

	if auth, err := os.LookupEnv("PAPERLESS_AUTH_TOKEN"); err {
		config.PAPERLESS_AUTH_TOKEN = auth
	} else {
		return nil, errors.New("PAPERLESS_AUTH_TOKEN not defined")
	}
	if val, err := os.LookupEnv("PAPERLESS_URL"); err {
		config.PAPERLESS_URL = val
	} else {
		return nil, errors.New("PAPERLESS_URL not defined")
	}

	if val, err := os.LookupEnv("FRONTEND_URL"); err {
		config.FRONTEND_URL = val
	} else {
		return nil, errors.New("FRONTEND_URL not defined")
	}
	if val, err := os.LookupEnv("REDIS_USER"); err {
		config.REDIS_USER = val
	} else {
		return nil, errors.New("REDIS_USER not defined")
	}
	if val, err := os.LookupEnv("REDIS_PASSWORD"); err {
		config.REDIS_PASSWORD = val
	} else {
		return nil, errors.New("REDIS_PASSWORD not defined")
	}
	if val, err := os.LookupEnv("REDIS_DB"); err {
		config.REDIS_DB = val
	} else {
		return nil, errors.New("REDIS_DB not defined")
	}
	if val, err := os.LookupEnv("REDIS_ADDRESS"); err {
		config.REDIS_ADDRESS = val
	} else {
		return nil, errors.New("REDIS_ADDRESS not defined")
	}

	if val, err := strconv.Atoi(getEnv("APP_PORT", "8000")); err == nil {
		config.APP_PORT = val
	} else {
		return nil, err
	}

	if val, err := strconv.ParseBool(getEnv("PAPERLESS_UNSAFE_SSL", "false")); err == nil {
		config.PAPERLESS_UNSAFE_SSL = val
	} else {
		return nil, err
	}

	if val, err := strconv.ParseBool(getEnv("DEBUG_MODE", "false")); err == nil {
		config.DEBUG = val
	} else {
		return nil, err
	}

	return &config, nil
}
