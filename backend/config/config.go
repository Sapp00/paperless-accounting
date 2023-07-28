package config

import (
	"errors"
	"log"
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
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New() (*Config, error) {

	err := godotenv.Load(".env")
	config := Config{
		PAPERLESS_EXPENSE_TAG: getEnv("PAPERLESS_EXPENSE_TAG", "expense"),
		PAPERLESS_INCOME_TAG:  getEnv("PAPERLESS_INCOME_TAG", "income"),
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
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
		config.PAPERLESS_UNSAFE_SSL = val
	} else {
		return nil, err
	}

	return &config, nil
}
