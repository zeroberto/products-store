package config

import (
	"testing"

	"github.com/zeroberto/products-store/products-list-endpoint/config"
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

func TestValidate_WhenGrpcClientConfigIsInvalid_ThenFailure(t *testing.T) {
	expectedMessage := "GrpcClientConfig is invalid"

	appConfig, _ := config.ReadConfig("applicationTest.yml")
	appConfig.GrpcClientConfig = *new(config.GrpcClientConfig)

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
