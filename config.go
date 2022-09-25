package main

import "os"

type AppConfig struct {
	appPort                  string
	driverType               string
	databaseConnectionString string
	rabbitmqConnectionString string
}

func getEnvVariableOrDefault(variableName string, defaultValue string) string {
	result := os.Getenv(variableName)
	if result != "" {
		return result
	}
	return defaultValue
}

func getConfig() AppConfig {
	return AppConfig{
		appPort:                  getEnvVariableOrDefault("PORT", "8080"),
		driverType:               getEnvVariableOrDefault("DRIVER_TYPE", "postgres"),
		databaseConnectionString: getEnvVariableOrDefault("DATABASE_CONNECTION_STRING", "postgresql://postgres:example@localhost:5432?sslmode=disable"),
		rabbitmqConnectionString: getEnvVariableOrDefault("RABBITMQ_CONNECTION_STRING", "amqp://guest:guest@localhost:5672/"),
	}
}
