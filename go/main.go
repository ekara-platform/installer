package main

import (
	"log"
	"os"

	"github.com/ekara-platform/engine/util"
	"github.com/ekara-platform/installer"
)

type NodeExtraVars struct {
	Params    map[string]string
	Instances int
}

// main starts the process of the installer.
//
// This method is supposed to be launched via an entrypoint through the Dockerfile
// used to generate the image.
//
func main() {
	c := installer.CreateContext(log.New(os.Stdout, util.InstallerLogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds))
	c.Log().Println(installer.LOG_STARTING)
	e := installer.Run(*c)
	if e != nil {
		c.Log().Fatal(e)
	}
}
