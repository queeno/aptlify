package cmd

import (
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/action"
	"github.com/queeno/aptlify/config"
)

func apply(cmd *commander.Command, args []string) error {

	var new_state config.ConfigStruct = config.ConfigStruct{}

	actions := action.CreateActions(context.Config(), context.State())

	for _, action := range actions {
		action.Apply(context.Config(), &new_state)
	}

	context.WriteState(new_state)

	return nil
}

func makeCmdApply() *commander.Command {
	return &commander.Command{
		Run:       apply,
		UsageLine: "apply",
		Short:     "Apply changes to aptly",
	}
}
