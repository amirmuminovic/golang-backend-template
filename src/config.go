package main

import "os"

type AppConfig struct {
	appPort          string
	driverType       string
	connectionString string
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
		appPort:          getEnvVariableOrDefault("PORT", "8080"),
		driverType:       getEnvVariableOrDefault("DRIVER_TYPE", "postgres"),
		connectionString: getEnvVariableOrDefault("CONNECTION_STRING", "postgresql://postgres:example@localhost:5432?sslmode=disable"),
	}
}
