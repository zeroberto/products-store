package appcontainer

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/zeroberto/products-store/discount-calculator/config"
	"github.com/zeroberto/products-store/discount-calculator/container"
)

// AppContainer is responsible for implementing the application container concepts
type AppContainer struct {
	mux       sync.RWMutex
	instances map[string]interface{}
	AppConfig *config.AppConfig
}

// Initialize is responsible for initializing all application structs
func (sc *AppContainer) Initialize(configFilename string) error {
	appConfig, err := config.ReadConfig(configFilename)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("did not read config file: file=%s, err=%v", configFilename, err))
	}

	if err := config.Validate(appConfig); err != nil {
		return errors.Wrap(err, "AppConfig is invalid")
	}

	sc.AppConfig = appConfig

	return nil
}

// Get is responsible for returning the instance of a struct
func (sc *AppContainer) Get(key string) (instance interface{}, ok bool) {
	sc.mux.RLock()
	instance, ok = sc.instances[key]
	sc.mux.RUnlock()

	return
}

// GetAppConfig is responsible for returning de application configuration properties
func (sc *AppContainer) GetAppConfig() *config.AppConfig {
	return sc.AppConfig
}

// Put is responsible for adding an instance of a struct to the container
func (sc *AppContainer) Put(key string, value interface{}) {
	sc.mux.Lock()
	sc.instances[key] = value
	sc.mux.Unlock()
}

// Remove is responsible for remove an instance of a struct to the container
func (sc *AppContainer) Remove(key string) {
	sc.mux.Lock()
	delete(sc.instances, key)
	sc.mux.Unlock()
}

// NewContainer is responsible for creating a new instance of Container
func NewContainer() container.Container {
	c := new(AppContainer)
	c.instances = make(map[string]interface{})
	return c
}
