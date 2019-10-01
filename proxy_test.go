package main

import (
	"github.com/ekara-platform/engine/action"
	"log"
	"os"
	"testing"

	"github.com/ekara-platform/engine/util"
	"github.com/stretchr/testify/assert"
)

func TestNoProxy(t *testing.T) {
	c := &installerContext{}
	c.logger = log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds)
	os.Setenv(util.ActionEnvVariableKey, action.ApplyActionID)
	os.Unsetenv(envHTTPProxy)
	os.Unsetenv(envHTTPSProxy)
	os.Unsetenv(envNoProxy)
	fillProxy(c)
	assert.Equal(t, "", c.proxy.Http)
	assert.Equal(t, "", c.proxy.Https)
	assert.Equal(t, "", c.proxy.NoProxy)
}

func TestProxy(t *testing.T) {
	c := &installerContext{}
	c.logger = log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds)
	os.Setenv(util.ActionEnvVariableKey, action.ApplyActionID)
	os.Setenv(envHTTPProxy, "http_value")
	os.Setenv(envHTTPSProxy, "https_value")
	os.Setenv(envNoProxy, "no_value")
	fillProxy(c)
	assert.Equal(t, "http_value", c.proxy.Http)
	assert.Equal(t, "https_value", c.proxy.Https)
	assert.Equal(t, "no_value", c.proxy.NoProxy)
}
