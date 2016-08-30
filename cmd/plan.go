package cmd

import (
	"github.com/gonuts/commander"
)

func plan(cmd *commander.Command, args []string) error {
	return nil
}

func makeCmdPlan() *commander.Command {
	cmd := &commander.Command{
		Run:        plan,
		UsageLine: "plan",
		Short:     "Generate a plan of execution",
	}
	return cmd
}
