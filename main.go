package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/ekara-platform/engine/util"
)

const ekaraWorkDir = "/var/lib/ekara"

// main starts the process of the installer.
//
// This method is supposed to be launched via an entry point through the Dockerfile
// used to generate the image.
//
func main() {
	logger := log.New(os.Stdout, util.InstallerLogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	context := createInstallerContext(logger)

	defer func() {
		if err := recover(); err != nil { //catch
			context.Feedback().Error("Panic: %v\n%s", err, string(debug.Stack()))
			os.Exit(2)
		}
	}()

	logger.Println(logStarting)
	e := Run(context)
	if e != nil {
		context.Feedback().Error(e.Error())
		os.Exit(1)
	}
}
