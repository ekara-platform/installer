package main

import (
	"fmt"
	"github.com/ekara-platform/engine/action"
	"log"
	"os"
	"path/filepath"

	"github.com/ekara-platform/model"

	"github.com/ekara-platform/engine"
	"github.com/ekara-platform/engine/ssh"
	"github.com/ekara-platform/engine/util"
)

const (
	envHTTPProxy  string = "http_proxy"
	envHTTPSProxy string = "https_proxy"
	envNoProxy    string = "no_proxy"
)

func Run(logger *log.Logger) (e error) {
	c := createInstallerContext(logger)

	// Fill launch context common properties
	e = fillContext(c)
	if e != nil {
		return e
	}

	a := os.Getenv(util.ActionEnvVariableKey)
	switch a {
	case action.ApplyActionID:
		c.Log().Println(LogActionApply)

		// Fill the SSH keys as they are needed for apply
		if e := fillSSHKeys(c); e != nil {
			return e
		}

		// Create the engine
		var ekara engine.Ekara
		ekara, e = engine.Create(c, ekaraWorkDir)
		if e != nil {
			return
		}

		// Initialize it to build the environment
		e = ekara.Init()
		if e != nil {
			return e
		}

		_, e = ekara.ActionManager().Run(action.ApplyActionID)
	default:
		if a == "" {
			a = LogNoAction
		}
		// Bad luck; unsupported action!
		e = fmt.Errorf(ErrorUnsupportedAction, a)
	}
	return
}

func fillContext(c *installerContext) error {
	fillProxy(c)
	if e := fillExchangeFolder(c); e != nil {
		return e
	}
	if e := fillLocation(c); e != nil {
		return e
	}
	if e := fillTemplateContext(c); e != nil {
		return e
	}
	return nil
}

// fillProxy loads the proxy settings form the environmant variables into the
// context
func fillProxy(c *installerContext) {
	c.proxy = model.Proxy{
		Http:    os.Getenv(envHTTPProxy),
		Https:   os.Getenv(envHTTPSProxy),
		NoProxy: os.Getenv(envNoProxy)}
}

func fillExchangeFolder(c *installerContext) error {
	var err error
	c.ef, err = util.CreateExchangeFolder(util.InstallerVolume, "")
	if err != nil {
		return fmt.Errorf("error creating the installer exchange folder: %s", err.Error())
	}
	return nil
}

// fillLocation extracts the descriptor location and descriptor file name from the
// environment variables:
//  - "engine.StarterEnvVariableKey"
//  - "engine.StarterEnvNameVariableKey"
//
// In addition it extract the user login to log into the repository where the
// descriptor is stored. this is done with the environment variables:
//  - "engine.StarterEnvLoginVariableKey"
//  - "engine.StarterEnvLoginVariableKey"
func fillLocation(c *installerContext) error {
	c.location = os.Getenv(util.StarterEnvVariableKey)
	if c.location == "" {
		return fmt.Errorf(ErrorRequiredEnv, util.StarterEnvVariableKey)
	}
	c.descriptorName = os.Getenv(util.StarterEnvNameVariableKey)
	if c.descriptorName == "" {
		return fmt.Errorf(ErrorRequiredEnv, util.StarterEnvNameVariableKey)
	}
	c.user = os.Getenv(util.StarterEnvLoginVariableKey)
	c.password = os.Getenv(util.StarterEnvPasswordVariableKey)
	return nil
}

func fillTemplateContext(c *installerContext) error {
	ok := c.Ef().Location.Contains(util.ExternalVarsFilename)
	if ok {
		var e error
		c.extVars, e = model.ParseParameters(util.JoinPaths(c.Ef().Location.Path(), util.ExternalVarsFilename))
		if e != nil {
			return fmt.Errorf(ErrorLoadingClientParameters, e)
		}
		c.Log().Printf(LogCLiParameters, c.extVars)
	}
	return nil
}

// fSHKeys checks if the SSH keys are specified via environment variables.
//
// If:
//		YES; they will be loaded into the context
//		NOT; they will be generated and then loaded into the context
//
func fillSSHKeys(c *installerContext) error {
	if c.Ef().Input.Contains(util.SSHPuplicKeyFileName) && c.Ef().Input.Contains(util.SSHPrivateKeyFileName) {
		c.sshPublicKey = filepath.Join(c.Ef().Input.Path(), util.SSHPuplicKeyFileName)
		c.sshPrivateKey = filepath.Join(c.Ef().Input.Path(), util.SSHPrivateKeyFileName)
		c.Log().Println("Using provided SSH keys")
	} else {
		c.Log().Println("Generating a new set of SSH keys")
		publicKey, privateKey, e := ssh.Generate()
		if e != nil {
			return fmt.Errorf(ErrorGeneratingSShKeys, e.Error())
		}
		_, e = util.SaveFile(c.Ef().Input, util.SSHPuplicKeyFileName, publicKey)
		if e != nil {
			return fmt.Errorf("an error occurred saving the public key into: %v", c.Ef().Input.Path())
		}
		_, e = util.SaveFile(c.Ef().Input, util.SSHPrivateKeyFileName, privateKey)
		if e != nil {
			return fmt.Errorf("an error occurred saving the private key into: %v", c.Ef().Input.Path())
		}
		c.sshPublicKey = filepath.Join(c.Ef().Input.Path(), util.SSHPuplicKeyFileName)
		c.sshPrivateKey = filepath.Join(c.Ef().Input.Path(), util.SSHPrivateKeyFileName)
	}

	if c.Log() != nil {
		c.Log().Printf(LogSSHPublicKey, c.sshPublicKey)
		c.Log().Printf(LogSSHPrivateKey, c.sshPrivateKey)
	}
	return nil
}
