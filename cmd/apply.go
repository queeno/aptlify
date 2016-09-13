package cmd

import (
	"github.com/queeno/aptlify/action"
	"github.com/queeno/aptlify/config"
	"github.com/smira/commander"
)

func apply(cmd *commander.Command, args []string) error {

	var newState config.ConfigStruct = config.ConfigStruct{}

	actions := action.CreateActions(context.Config(), context.State())

	for _, action := range actions {
		action.Apply(context.Config(), &newState)
	}

	context.WriteState(newState)

	return nil
}

func makeCmdApply() *commander.Command {
	return &commander.Command{
		Run:       apply,
		UsageLine: "apply",
		Short:     "Apply changes to aptly",
	}
}
