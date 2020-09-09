package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// AppConfig reflects the all application properties
type AppConfig struct {
	DSConfig     DSConfig      `yaml:"dataStore"`
	ServerConfig NetworkConfig `yaml:"server"`
}

// DSConfig reflects the data store config properties
type DSConfig struct {
	UserInfoConfig DBConfig `yaml:"userInfo"`
}

// DBConfig reflects the database config properties
type DBConfig struct {
	Type         string     `yaml:"type"`
	Host         string     `yaml:"host"`
	Port         uint       `yaml:"port"`
	Database     string     `yaml:"database"`
	DatabaseType string     `yaml:"databaseType"`
	AuthConfig   AuthConfig `yaml:"auth"`
}

// NetworkConfig reflects the general network config properties
type NetworkConfig struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}

// AuthConfig reflects the general authentication properties
type AuthConfig struct {
	Type string `yaml:"type"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Repo string `yaml:"repo"`
}

// GetURI returns the complete URL with concatenated host and port
func (dbc *DBConfig) GetURI() string {
	return fmt.Sprintf("%s:%d", dbc.Host, dbc.Port)
}

// ReadConfig is responsible for read the config file
func ReadConfig(configFileName string) (*AppConfig, error) {
	var appConfig AppConfig
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, errors.Wrap(err, "An error occurred while reading the file")
	}

	err = yaml.Unmarshal(file, &appConfig)
	if err != nil {
		return nil, errors.Wrap(err, "An error occurred while deserializing the file")
	}

	return &appConfig, nil
}

// Validate is responsible for verifying that an appConfig instance is valid
func Validate(ac *AppConfig) error {
	emptyAC := AppConfig{}

	if *ac == emptyAC {
		return errors.New("AppConfig is invalid")
	}
	if dsc := new(DSConfig); ac.DSConfig == *dsc {
		return errors.New("DSConfig is invalid")
	}
	if !validateDBConfig(ac.DSConfig.UserInfoConfig) {
		return errors.New("UserInfoConfig is invalid")
	}
	if sc := new(NetworkConfig); ac.ServerConfig == *sc {
		return errors.New("ServerConfig is invalid")
	}
	return nil
}

func validateDBConfig(dbc DBConfig) bool {
	if dbc.Type != "" && dbc.Host != "" && dbc.Port != 0 && dbc.Database != "" && dbc.DatabaseType != "" {
		return true
	}
	return false
}
