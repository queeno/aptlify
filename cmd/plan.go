package cmd

import (
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/repo"
	"github.com/queeno/aptlify/gpg"
	"github.com/queeno/aptlify/action"
)



func createActions() []action.ActionStruct {

	config := context.Config()
	state := context.State()

	mirrorActions := mirror.CreateMirrorActions(config.Mirrors, state.Mirrors)
	repoActions := repo.CreateRepoActions(config.Repos, state.Repos)
	gpgActions := gpg.CreateGpgActions(config.Gpg_keys, state.Gpg_keys)

	return append(append(mirrorActions, repoActions...), gpgActions...)
}

func plan(cmd *commander.Command, args []string) error {

	// Create changes

	actions := createActions()

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
