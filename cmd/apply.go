package cmd

import (
	"github.com/gonuts/commander"
)

func apply(cmd *commander.Command, args []string) error {

return nil


}

func makeCmdApply() *commander.Command {
	return &commander.Command{
		Run:       apply,
		UsageLine: "apply",
		Short:     "Apply changes to aptly",
	}
}
