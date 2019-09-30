package main

import (
	"github.com/ekara-platform/engine/action"
	"log"
	"os"
	"testing"

	"github.com/ekara-platform/engine/util"
	"github.com/stretchr/testify/assert"
)

func TestNoLocation(t *testing.T) {
	c := installerContext{}
	c.logger = log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds)
	os.Setenv(util.ActionEnvVariableKey, action.ApplyActionID)
	os.Unsetenv(util.StarterEnvVariableKey)
	os.Setenv(util.StarterEnvNameVariableKey, "test_name")
	e := fillLocation(&c)
	assert.NotNil(t, e)
	assert.Equal(t, e.Error(), "the environment variable \"EKARA_ENV_DESCR\" should be defined")
}

func TestNoName(t *testing.T) {
	c := installerContext{}
	c.logger = log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds)
	os.Setenv(util.ActionEnvVariableKey, action.ApplyActionID)
	os.Unsetenv(util.StarterEnvNameVariableKey)
	os.Setenv(util.StarterEnvVariableKey, "test_location")
	e := fillLocation(&c)
	assert.NotNil(t, e)
	assert.Equal(t, e.Error(), "the environment variable \"EKARA_ENV_DESCR_NAME\" should be defined")
}

func TestLocation(t *testing.T) {
	c := installerContext{}
	c.logger = log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds)
	os.Setenv(util.ActionEnvVariableKey, action.ApplyActionID)
	os.Setenv(util.StarterEnvVariableKey, "test_location")
	os.Setenv(util.StarterEnvNameVariableKey, "test_name")
	e := fillLocation(&c)
	assert.Nil(t, e)
	assert.Equal(t, c.location, "test_location")
	assert.Equal(t, c.descriptorName, "test_name")
}
