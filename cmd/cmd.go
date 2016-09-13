package cmd

import (
	"github.com/smira/commander"
	"github.com/smira/flag"
	"os"
)

func LookupOption(defaultValue bool, flags *flag.FlagSet, name string) (result bool) {
	result = defaultValue
	if flags.IsSet(name) {
		result = flags.Lookup(name).Value.Get().(bool)
	}
	return
}

func RootCommand() *commander.Command {

	cmd := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "Manage aptly with aptlify",
		Long: `aptlify manages the publishing of aptly repos and mirrors into a
unique, simple, configuration-driven tool`,
		Flag: *flag.NewFlagSet("aptlify", flag.ExitOnError),
		Subcommands: []*commander.Command{
			makeCmdApply(),
			makeCmdPlan(),
			makeCmdDump(),
			makeCmdMirror(),
		},
	}

	cmd.Flag.String("config", "", "location of configuration file (default location is ~/.aptlify.conf")

	return cmd

}
