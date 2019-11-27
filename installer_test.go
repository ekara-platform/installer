package main

import (
	"fmt"
	"github.com/ekara-platform/engine/action"
	"log"
	"os"
	"testing"

	"github.com/ekara-platform/engine/util"
	"github.com/stretchr/testify/assert"
)

/*
ActionFailId ActionId 1
ActionReportId 2
ActionCreateId 3
ActionInstallId 4
ActionDeployId 5
ActionCheckId 6
ActionDumpId 7
ActionUpdateId 8
ActionDeleteId 9
ActionNilId 10

case engine.ActionCreateId.String(): 3
case engine.ActionInstallId.String(): 4
case engine.ActionDeployId.String(): 5
case engine.ActionCheckId.String(): 6
case engine.ActionDumpId.String():  7
*/
func TestNoAction(t *testing.T) {
	os.Unsetenv(util.ActionEnvVariableKey)
	os.Setenv(util.StarterEnvVariableKey, "DummyDescriptor")
	os.Setenv(util.StarterEnvNameVariableKey, "DummyDescriptorName")

	e := Run(createInstallerContext(log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime|log.Lmicroseconds)))
	assert.NotNil(t, e)
	assert.Equal(t, "the action \"No action specified\" is not supported by the installer", e.Error())
}

func checkUnsupportedAction(t *testing.T, a action.ActionID) {
	os.Setenv(util.ActionEnvVariableKey, a.String())
	e := Run(createInstallerContext(log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime|log.Lmicroseconds)))
	assert.NotNil(t, e)
	assert.Equal(t, e.Error(), fmt.Sprintf("the action \"%s\" is not supported by the installer", a))
}

func TestWrongActions(t *testing.T) {
	os.Setenv(util.StarterEnvVariableKey, "DummyDescriptor")
	os.Setenv(util.StarterEnvNameVariableKey, "DummyDescriptorName")

	checkUnsupportedAction(t, action.CheckActionID)
	checkUnsupportedAction(t, action.DumpActionID)
	checkUnsupportedAction(t, action.NilActionID)
}
