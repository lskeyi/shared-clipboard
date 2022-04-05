package commands

import (
	"fmt"
	"github.com/jutkko/cli"
	"shared-clipboard/runcommand"
)

// TargetCommand is the command that is responsible setting the copy-pasta
// targets
type TargetCommand struct {
	Ui cli.Ui
}

// Help string
func (t *TargetCommand) Help() string {
	return `Usage: copy-pasta target [<target>]

    Changes the current target to the target.
    If no argument is provided, it lists the current target.
`
}

// Run function for the command
func (t *TargetCommand) Run(args []string) int {
	config, err := loadRunCommands()
	if err != nil {
		return 1
	}

	if len(args) > 0 {
		if target, ok := config.Targets[args[0]]; ok {
			if err := runcommand.Update(target.Name, target.SecretId, target.SecretKey, target.BucketName, target.Location); err != nil {
				t.Ui.Error(fmt.Sprintf("Failed to update the current target: %s", err.Error()))
				return 2
			} else {
				return 0
			}
		} else {
			t.Ui.Error("Target is invalid")
			return 3
		}
	} else {
		t.Ui.Output("copy-pasta current target:")
		t.Ui.Output("  " + config.CurrentTarget.Name)
		return 0
	}
}

// Synopsis is the short help string
func (t *TargetCommand) Synopsis() string {
	return "Changes the current target to the provided target"
}
