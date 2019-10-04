package main

import (
	"log"
	"os"

	"github.com/ekara-platform/engine/util"
)

const ekaraWorkDir = "/var/lib/ekara"

// main starts the process of the installer.
//
// This method is supposed to be launched via an entrypoint through the Dockerfile
// used to generate the image.
//
func main() {
	logger := log.New(os.Stdout, util.InstallerLogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	logger.Println(logStarting)
	e := Run(logger)
	if e != nil {
		logger.Fatal(e)
	}
}
