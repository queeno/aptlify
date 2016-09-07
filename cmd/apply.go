package cmd

import (
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/action"
)

func apply(cmd *commander.Command, args []string) error {

	actions := action.CreateActions(context.Config(), context.State())

	for _, action := range actions {
		action.Apply(context.Config())
	}
	return nil
}

func makeCmdApply() *commander.Command {
	return &commander.Command{
		Run:       apply,
		UsageLine: "apply",
		Short:     "Apply changes to aptly",
	}
}
