package main

import (
	"fmt"
	"os"

	"github.com/ekara-platform/engine/action"

	"github.com/ekara-platform/engine"
	"github.com/ekara-platform/engine/util"
)

//Run starts the installer
func Run(c *installerContext) (error) {
	// Fill launch context common properties
	e := fillContext(c)
	if e != nil {
		return e
	}

	var result action.Result
	a := os.Getenv(util.ActionEnvVariableKey)
	switch a {
	case action.ApplyActionID:
		c.Log().Println(logActionApply)
		// Fill the SSH keys as they are needed for apply and destroy
		if e := fillSSHKeys(c); e != nil {
			return e
		}
		result, e = executeAction(c, action.ApplyActionID)
		if e != nil {
			return e
		}
	case action.DestroyActionID:
		c.Log().Println(logActionDestroy)
		// Fill the SSH keys as they may be needed for destroy
		if e := fillSSHKeys(c); e != nil {
			return e
		}
		result, e = executeAction(c, action.DestroyActionID)
		if e != nil {
			return e
		}
	default:
		if a == "" {
			a = logNoAction
		}
		// Bad luck; unsupported action!
		e = fmt.Errorf(errorUnsupportedAction, a)
	}

	if result != nil {
		str, e := result.AsJson()
		if e != nil {
			return e
		}
		var path string
		path, e = util.SaveFile(c.Ef().Output, "result.json", []byte(str))
		if e != nil {
			return e
		}
		c.logger.Printf("Action result written in %s\n", path)
	}
	return nil
}

func executeAction(c *installerContext, action action.ActionID) (action.Result, error) {
	// Create the engine
	var ekara engine.Ekara
	var e error
	ekara, e = engine.Create(c, ekaraWorkDir)
	if e != nil {
		return nil, e
	}
	// Initialize it to build the environment
	e = ekara.Init()
	if e != nil {
		return nil, e
	}
	// Execute the action
	return ekara.ActionManager().Run(action)
}
