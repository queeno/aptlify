package cmd

import (
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/action"
	"github.com/queeno/aptlify/config"
)

func mirror(cmd *commander.Command, args []string) error {

	var newState config.ConfigStruct = config.ConfigStruct{}

	mirrors := context.State().Mirrors
	actions := action.UpdateMirrors(mirrors)

	for _, action := range actions {
		action.Apply(context.State(), &newState)
	}

	return nil
}

func makeCmdMirror() *commander.Command {
	return &commander.Command{
		Run:       mirror,
		UsageLine: "mirror",
		Short:     "Update mirrors",
	}
}
