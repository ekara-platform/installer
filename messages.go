package main

const (
	ErrorRequiredEnv             string = "the environment variable \"%s\" should be defined"
	ErrorLoadingClientParameters string = "error loading the CLI parameters : %s"
	ErrorGeneratingSShKeys       string = "error generating the SSH keys %s"
	ErrorUnsupportedAction       string = "the action \"%s\" is not supported by the installer"

	LogStarting string = "Starting the installer..."
	LogRunning  string = "Running the installer..."

	LogActionApply string = "Apply action requested"
	LogNoAction    string = "No action specified"

	LogSSHPublicKey  string = "Installer using SSH public key: %s"
	LogSSHPrivateKey string = "Installer using SSH private key: %s"

	LogCLiParameters string = "Using CLI parameters: %v"
)
