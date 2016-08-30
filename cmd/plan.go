package cmd

import (
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/gpg"
)

func plan(cmd *commander.Command, args []string) error {
	var err error
	err = gpg.Plan()
	if err != nil {
		return err
	}
	// err = publish.Plan()
	// if err != nil {
	//    return err
	//}
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
