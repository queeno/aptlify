package cmd

import (
	"github.com/gonuts/commander"
	"os"
)

func RootCommand() *commander.Command {

	cmd := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "Manage aptly with aptlify",
		Long: `aptlify manages the publishing of aptly repos and mirrors into a
unique, simple, configuration-driven tool`,
		Subcommands: []*commander.Command{
			makeCmdApply(),
			makeCmdPlan(),
			makeCmdDump(),
			makeCmdMirror(),
		},
	}

	return cmd

}
