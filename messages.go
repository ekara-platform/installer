package main

const (
	errorRequiredEnv             string = "the environment variable \"%s\" should be defined"
	errorLoadingClientParameters string = "error loading the CLI parameters : %s"
	errorGeneratingSShKeys       string = "error generating the SSH keys %s"
	errorUnsupportedAction       string = "the action \"%s\" is not supported by the installer"

	logStarting string = "Starting the installer..."
	logRunning  string = "Running the installer..."

	logActionApply   string = "Apply action requested"
	logActionDestroy string = "Destroy action requested"
	logNoAction      string = "No action specified"

	logSSHPublicKey  string = "Installer using SSH public key: %s"
	logSSHPrivateKey string = "Installer using SSH private key: %s"

	logCLiParameters string = "Using CLI parameters: %v"
)
