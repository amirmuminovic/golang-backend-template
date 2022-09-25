package main

import (
	"os"
	"testing"
)

func TestGetDefaultConfig(t *testing.T) {
	ac := getConfig()
	defaultConfig := AppConfig{
		driverType:       "postgres",
		connectionString: "postgresql://postgres:example@localhost:5432?sslmode=disable",
		appPort:          "8080",
	}

	if ac.appPort != defaultConfig.appPort {
		t.Errorf("Default config for the appPort changed. Expected %s but received %s", defaultConfig.appPort, ac.appPort)
	}
	if ac.connectionString != defaultConfig.connectionString {
		t.Errorf("Default config for the connectionString changed. Expected %s but received %s", defaultConfig.connectionString, ac.connectionString)
	}
	if ac.driverType != defaultConfig.driverType {
		t.Errorf("Default config for the driverType changed. Expected %s but received %s", defaultConfig.driverType, ac.driverType)
	}
}

func TestGetConfig(t *testing.T) {
	customPort := "3000"
	customDriverType := "mysql"
	customConnectionString := "connectionString"
	os.Setenv("PORT", customPort)
	os.Setenv("DRIVER_TYPE", customDriverType)
	os.Setenv("CONNECTION_STRING", customConnectionString)
	ac := getConfig()

	if ac.appPort != customPort {
		t.Errorf("Default config for the appPort changed. Expected %s but received %s", customPort, ac.appPort)
	}
	if ac.connectionString != customConnectionString {
		t.Errorf("Default config for the connectionString changed. Expected %s but received %s", customConnectionString, ac.connectionString)
	}
	if ac.driverType != customDriverType {
		t.Errorf("Default config for the driverType changed. Expected %s but received %s", customDriverType, ac.driverType)
	}

	os.Clearenv()
}
