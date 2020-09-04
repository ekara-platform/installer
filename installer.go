package main

import (
    "fmt"
    "github.com/GroupePSA/componentizer"
    "os"
    "strings"

    "github.com/ekara-platform/engine/action"

    "github.com/ekara-platform/engine"
    "github.com/ekara-platform/engine/util"
)

//Run starts the installer
func Run(c *installerContext) error {
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
        return fmt.Errorf(errorUnsupportedAction, a)
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
    repo, e := buildRepositoryFromEnv()
    if e != nil {
        return nil, e
    }

    // Create the engine
    ekara := engine.Create(c, ekaraWorkDir)
    e = ekara.Init(repo)
    if e != nil {
        return nil, e
    }

    // Execute the action
    return ekara.Execute(action)
}

func buildRepositoryFromEnv() (componentizer.Repository, error) {
    // Repository location
    location := os.Getenv(util.StarterEnvVariableKey)
    ref := "master"
    if location == "" {
        return componentizer.Repository{}, fmt.Errorf(errorRequiredEnv, util.StarterEnvVariableKey)
    }
    split := strings.Split(location, "@")
    if len(split) > 1 {
        location = split[0]
        ref = split[1]
    }

    // Optional repository authentication
    user := os.Getenv(util.StarterEnvLoginVariableKey)
    password := os.Getenv(util.StarterEnvPasswordVariableKey)
    auth := make(map[string]string)
    if user != "" || password != "" {
        auth["method"] = "basic"
        auth["user"] = user
        auth["password"] = password
    }

    return componentizer.CreateRepository(location, ref, auth)
}
