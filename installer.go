package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ekara-platform/engine/action"

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

//Run starts the installer
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
		c.Log().Println(logActionApply)

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

		applyResult, e := ekara.ActionManager().Run(action.ApplyActionID)
		if e != nil {
			return e
		}
		logger.Println("----------------------------------------------------------------------")
		logger.Println("Inventory")
		logger.Println("----------------------------------------------------------------------")
		for _, host := range applyResult.(action.ApplyResult).Inventory.Hosts {
			logger.Println(host.Name)
		}
		logger.Println("----------------------------------------------------------------------")
	default:
		if a == "" {
			a = logNoAction
		}
		// Bad luck; unsupported action!
		e = fmt.Errorf(errorUnsupportedAction, a)
	}
	return
}

func fillContext(c *installerContext) error {
	fillProxy(c)
	fillVerbosity(c)
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

// fillProxy loads the proxy settings form the environment variables into the
// context
func fillProxy(c *installerContext) {
	c.proxy = model.Proxy{
		Http:    os.Getenv(envHTTPProxy),
		Https:   os.Getenv(envHTTPSProxy),
		NoProxy: os.Getenv(envNoProxy)}
}

// fillVerbosity fills the engine verbosity level based on an environment variable
func fillVerbosity(c *installerContext) {
	var err error
	c.verbosity, err = strconv.Atoi(os.Getenv(util.StarterVerbosityVariableKey))
	if err != nil {
		c.verbosity = 2
	}
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
		return fmt.Errorf(errorRequiredEnv, util.StarterEnvVariableKey)
	}
	c.descriptorName = os.Getenv(util.StarterEnvNameVariableKey)
	if c.descriptorName == "" {
		return fmt.Errorf(errorRequiredEnv, util.StarterEnvNameVariableKey)
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
			return fmt.Errorf(errorLoadingClientParameters, e)
		}
		c.Log().Printf(logCLiParameters, c.extVars)
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
	if c.Ef().Input.Contains(util.SSHPublicKeyFileName) && c.Ef().Input.Contains(util.SSHPrivateKeyFileName) {
		c.sshPublicKey = filepath.Join(c.Ef().Input.Path(), util.SSHPublicKeyFileName)
		c.sshPrivateKey = filepath.Join(c.Ef().Input.Path(), util.SSHPrivateKeyFileName)
		c.Log().Println("Using provided SSH keys")
	} else {
		c.Log().Println("Generating a new set of SSH keys")
		publicKey, privateKey, e := ssh.Generate()
		if e != nil {
			return fmt.Errorf(errorGeneratingSShKeys, e.Error())
		}
		_, e = util.SaveFile(c.Ef().Input, util.SSHPublicKeyFileName, publicKey)
		if e != nil {
			return fmt.Errorf("an error occurred saving the public key into: %v", c.Ef().Input.Path())
		}
		_, e = util.SaveFile(c.Ef().Input, util.SSHPrivateKeyFileName, privateKey)
		if e != nil {
			return fmt.Errorf("an error occurred saving the private key into: %v", c.Ef().Input.Path())
		}
		c.sshPublicKey = filepath.Join(c.Ef().Input.Path(), util.SSHPublicKeyFileName)
		c.sshPrivateKey = filepath.Join(c.Ef().Input.Path(), util.SSHPrivateKeyFileName)
	}

	if c.Log() != nil {
		c.Log().Printf(logSSHPublicKey, c.sshPublicKey)
		c.Log().Printf(logSSHPrivateKey, c.sshPrivateKey)
	}
	return nil
}
