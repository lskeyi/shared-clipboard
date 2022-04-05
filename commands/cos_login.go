package commands

import (
	"flag"
	"fmt"
	"github.com/jutkko/cli"
	"shared-clipboard/runcommand"
)

// S3LoginCommand is the command that is responsible for logging in, the
// size effect is that it saves the config file locally
type CosLoginCommand struct {
	Ui cli.Ui
}

// Help string
func (l *CosLoginCommand) Help() string {
	return `Usage: copy-pasta s3-login [--target] [<target>] [--endpoint] [<endpoint>] [--location] [<location>]

		Prompts to login interactively. The command expects S3 credentials. If no
		target is provided, the "default" target name is provided.

Options:
    --target       Specify the new target name.
    --endpoint     Specify the new target's endpoint, defaults to s3.amazonaws.com.
    --location     Specify the new target's location, defaults to eu-west-2.
`
}

// Run function for the command
func (l *CosLoginCommand) Run(args []string) int {
	loginCommand := flag.NewFlagSet("login", flag.ExitOnError)
	loginTargetOption := loginCommand.String("target", "default", "the name for copy-pasta's target")
	loginLocationOption := loginCommand.String("location", "eu-east-8", "the location for the backend bucket")

	// not tested, may be too hard
	err := loginCommand.Parse(args)
	if err != nil {
		l.Ui.Error(err.Error())
		return 10
	}

	accessKey, err := l.Ui.Ask("Please enter secretId:")
	if err != nil {
		l.Ui.Error(err.Error())
		return 10
	}

	secretAccessKey, err := l.Ui.AskSecret("Please enter secret key:")
	if err != nil {
		l.Ui.Error(err.Error())
		return 10
	}

	if err := runcommand.Update(*loginTargetOption, accessKey, secretAccessKey, "default-bucket-1306768814", *loginLocationOption); err != nil {
		l.Ui.Error(fmt.Sprintf("Failed to update the current target: %s\n", err.Error()))
		return 9
	}

	fmt.Println("Log in information saved")

	return 0
}

// Synopsis is the short help string
func (l *CosLoginCommand) Synopsis() string {
	return fmt.Sprintf("Login to copy-pasta with Cos backend")
}
