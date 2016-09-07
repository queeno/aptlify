package cmd

import (
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/action"
)

func plan(cmd *commander.Command, args []string) error {

	// Create changes

	actions := action.CreateActions(context.Config(), context.State())

	for _, action := range actions {
		action.Plan()
	}

	return nil
}

func makeCmdPlan() *commander.Command {
	cmd := &commander.Command{
		Run:       plan,
		UsageLine: "plan",
		Short:     "Generate a plan of execution",
	}
	return cmd
}
