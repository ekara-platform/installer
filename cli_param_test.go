package main

import (
	"log"
	"os"
	"testing"

	"github.com/ekara-platform/engine/util"
	"github.com/stretchr/testify/assert"
)

func TestReadingParam(t *testing.T) {
	ef, e := util.CreateExchangeFolder("./", "testFolder")
	assert.Nil(t, e)
	assert.NotNil(t, ef)
	defer ef.Delete()

	e = ef.Create()
	assert.Nil(t, e)

	pContent := `key1: value1`

	e = ef.Location.Write([]byte(pContent), util.ExternalVarsFilename)
	assert.Nil(t, e)

	c := &installerContext{
		ef:     ef,
		logger: log.New(os.Stdout, "Test", log.Ldate|log.Ltime|log.Lmicroseconds),
	}

	e = fillTemplateContext(c)
	assert.Nil(t, e)
	cParam := c.extVars
	assert.NotNil(t, cParam)
	assert.Equal(t, len(cParam), 1)
}
