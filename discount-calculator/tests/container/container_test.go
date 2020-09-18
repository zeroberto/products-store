package container

import (
	"strings"
	"testing"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/chrono/provider"
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/appcontainer"
)

func TestInitialize(t *testing.T) {
	var container container.Container = &appcontainer.AppContainer{}

	if err := container.Initialize("applicationTest.yml"); err != nil {
		t.Errorf("Initialize() failed, expected %v, got %v", nil, err)
	}
}

func TestInitialize_WhenConfigFileNotFound_ThenFailure(t *testing.T) {
	expectedMessage := "did not read config file: file=applicationNotFound.yml"
	var container container.Container = &appcontainer.AppContainer{}

	err := container.Initialize("applicationNotFound.yml")
	if err == nil {
		t.Error("Initialize() failed, expected an error, got nil")
	}
	if strings.Contains(expectedMessage, err.Error()) {
		t.Errorf("Initialize() failed, expected %v, got %v", expectedMessage, err.Error())
	}
}

func TestInitialize_WhenConfigFileIsInvalid_ThenFailure(t *testing.T) {
	expectedMessage := "AppConfig is invalid"
	var container container.Container = &appcontainer.AppContainer{}

	err := container.Initialize("applicationIncomplete.yml")
	if err == nil {
		t.Error("Initialize() failed, expected an error, got nil")
	}
	if strings.Contains(expectedMessage, err.Error()) {
		t.Errorf("Initialize() failed, expected %v, got %v", expectedMessage, err.Error())
	}
}

func TestPutAndGet(t *testing.T) {
	targetKey := "key"
	var expected chrono.TimeStamp = &provider.TimeStampImpl{}

	var container container.Container = &appcontainer.AppContainer{}
	container.Initialize("applicationTest.yml")
	container.Put(targetKey, expected)

	got, ok := container.Get(targetKey)
	if !ok {
		t.Errorf("Put() and Get() failed, expected at least one instance with the key %s", targetKey)
	}
	if got == nil {
		t.Error("Put() and Get() failed, expected an instance, got nil")
	}
	if got != expected {
		t.Errorf("Put() and Get() failed, expected %v, got %v", expected, got)
	}
}

func TestRemove(t *testing.T) {
	targetKey := "key"
	var expected chrono.TimeStamp = &provider.TimeStampImpl{}

	var container container.Container = &appcontainer.AppContainer{}
	container.Initialize("applicationTest.yml")
	container.Put(targetKey, expected)

	container.Remove(targetKey)

	got, ok := container.Get(targetKey)
	if ok {
		t.Errorf("Put() and Get() failed, no instance is expected with the key %s", targetKey)
	}
	if got != nil {
		t.Errorf("Put() and Get() failed, no instance is expected, got %v", got)
	}
}
