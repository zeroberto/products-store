package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// AppConfig reflects the all application properties
type AppConfig struct {
	GrpcClientConfig GrpcClientConfig `yaml:"grpcClientConfig"`
	ServerConfig     NetworkConfig    `yaml:"serverConfig"`
}

// GrpcClientConfig reflects the gRPC client config properties
type GrpcClientConfig struct {
	ProductsListConfig NetworkConfig `yaml:"productsListConfig"`
}

// NetworkConfig reflects the general network config properties
type NetworkConfig struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}

// GetURI returns the complete URL with concatenated host and port
func (nc *NetworkConfig) GetURI() string {
	return fmt.Sprintf("%s:%d", nc.Host, nc.Port)
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
	if gcc := new(GrpcClientConfig); ac.GrpcClientConfig == *gcc {
		return errors.New("GrpcClientConfig is invalid")
	}
	if sc := new(NetworkConfig); ac.ServerConfig == *sc {
		return errors.New("ServerConfig is invalid")
	}
	return nil
}
