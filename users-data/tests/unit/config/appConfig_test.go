package config

import (
	"testing"

	"github.com/zeroberto/products-store/users-data/config"
)

func TestReadConfig(t *testing.T) {
	if _, err := config.ReadConfig("applicationTest.yml"); err != nil {
		t.Errorf("ReadConfig() failed, expected %v, got %v", nil, err)
	}
}

func TestReadConfig_WhenFileNotExists_ThenFailure(t *testing.T) {
	if _, err := config.ReadConfig("applicationNotExists.yml"); err == nil {
		t.Error("ReadConfig() failed, expected an error, got nil")
	}
}

func TestValidate(t *testing.T) {
	appConfig, _ := config.ReadConfig("applicationTest.yml")

	if err := config.Validate(appConfig); err != nil {
		t.Errorf("ReadConfig() failed, expected %v, got %v", nil, err)
	}
}

func TestValidate_WhenConfigIsInvalid_ThenFailure(t *testing.T) {
	expectedMessage := "AppConfig is invalid"
	appConfig, _ := config.ReadConfig("applicationInvalid.yml")

	err := config.Validate(appConfig)
	if err == nil {
		t.Error("Validate() failed, expected an error, got nil")
	}
	if err.Error() != expectedMessage {
		t.Errorf("Validate() failed, expected %v, got %v", expectedMessage, err.Error())
	}
}

func TestValidate_WhenDSConfigIsInvalid_ThenFailure(t *testing.T) {
	expectedMessage := "DSConfig is invalid"

	appConfig, _ := config.ReadConfig("applicationTest.yml")
	appConfig.DSConfig = *new(config.DSConfig)

	err := config.Validate(appConfig)
	if err == nil {
		t.Error("Validate() failed, expected an error, got nil")
	}
	if err.Error() != expectedMessage {
		t.Errorf("Validate() failed, expected %v, got %v", expectedMessage, err.Error())
	}
}

func TestValidate_WhenUserInfoConfigIsInvalid_ThenFailure(t *testing.T) {
	expectedMessage := "UserInfoConfig is invalid"

	appConfig, _ := config.ReadConfig("applicationTest.yml")
	appConfig.DSConfig.UserInfoConfig.Port = 0

	err := config.Validate(appConfig)
	if err == nil {
		t.Error("Validate() failed, expected an error, got nil")
	}
	if err.Error() != expectedMessage {
		t.Errorf("Validate() failed, expected %v, got %v", expectedMessage, err.Error())
	}
}

func TestValidate_WhenServerConfigIsInvalid_ThenFailure(t *testing.T) {
	expectedMessage := "ServerConfig is invalid"

	appConfig, _ := config.ReadConfig("applicationTest.yml")
	appConfig.ServerConfig = *new(config.NetworkConfig)

	err := config.Validate(appConfig)
	if err == nil {
		t.Error("Validate() failed, expected an error, got nil")
	}
	if err.Error() != expectedMessage {
		t.Errorf("Validate() failed, expected %v, got %v", expectedMessage, err.Error())
	}
}
