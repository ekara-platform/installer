package main

import (
	"log"

	"github.com/ekara-platform/engine/util"
	"github.com/ekara-platform/model"
)

type (
	//InstallerContext Represents the informations required start the ekara engine
	// from the installer container
	installerContext struct {
		logger         *log.Logger
		ef             util.ExchangeFolder
		descriptorName string
		location       string
		user           string
		password       string
		sshPublicKey   string
		sshPrivateKey  string
		proxy          model.Proxy
		extVars        model.Parameters
	}
)

//DescriptorName implements the corresponding method in LaunchContext
func (c installerContext) DescriptorName() string {
	return c.descriptorName
}

//Location specifies where to look for the environment descriptor
func (c installerContext) Location() string {
	return c.location
}

//User The user to log into the descriptor repository
func (c installerContext) User() string {
	return c.user
}

//Password The user to log into the descriptor repository
func (c installerContext) Password() string {
	return c.password
}

//Log the looger to used during the ekara execution
func (c installerContext) Log() *log.Logger {
	return c.logger
}

//Ef the exchange folder where to find informations required
// to complete the Ekara execution of where to write informations
// produced by the execution.
// It will be a volume passed to the container by the CLI.
func (c installerContext) Ef() util.ExchangeFolder {
	return c.ef
}

//Proxy is the proxy info used by the engine during the process execution
func (c installerContext) Proxy() model.Proxy {
	return c.proxy
}

//SSHPublicKey the public key used by the engine during the process execution to
// connect the created nodes
func (c installerContext) SSHPublicKey() string {
	return c.sshPublicKey
}

//SSHPrivateKey the private key used by the engine during the process execution to
// connect the created nodes
func (c installerContext) SSHPrivateKey() string {
	return c.sshPrivateKey
}

//ExternalVars returns the external variables passed to the installer through the CLI
func (lC installerContext) ExternalVars() model.Parameters {
	return lC.extVars
}

//CreateContext returns a new installer context used to run the engine
func createInstallerContext(l *log.Logger) *installerContext {
	c := &installerContext{}
	c.logger = l
	return c
}
