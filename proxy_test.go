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
	os.Unsetenv("http_proxy")
	os.Unsetenv("https_proxy")
	os.Unsetenv("no_proxy")
	fillProxy(c)
	assert.Equal(t, "", c.proxy.Http)
	assert.Equal(t, "", c.proxy.Https)
	assert.Equal(t, "", c.proxy.NoProxy)
}

func TestProxy(t *testing.T) {
	c := &installerContext{}
	c.logger = log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds)
	os.Setenv(util.ActionEnvVariableKey, action.ApplyActionID)
	os.Setenv("http_proxy", "http_value")
	os.Setenv("https_proxy", "https_value")
	os.Setenv("no_proxy", "no_value")
	fillProxy(c)
	assert.Equal(t, "http_value", c.proxy.Http)
	assert.Equal(t, "https_value", c.proxy.Https)
	assert.Equal(t, "no_value", c.proxy.NoProxy)
}
