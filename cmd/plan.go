package cmd

import (
	"github.com/gonuts/commander"
	"fmt"
)

func plan(cmd *commander.Command, args []string) error {
	//var err error
	//err = gpg.Plan()
	//if err != nil {
	//	return err
	//}
	// err = publish.Plan()
	// if err != nil {
	//    return err
	//}

	fmt.Println(context.Config())

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
