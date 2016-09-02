package cmd

import (
	"fmt"
	"github.com/gonuts/commander"
	"github.com/queeno/aptlify/aptly"
	ctx "github.com/queeno/aptlify/context"
	"github.com/queeno/aptlify/utils"
)

func dump(cmd *commander.Command, args []string) error {

	ctx.Logging.Info.Println("retrieving information from aptly...")

	a := &aptly.AptlyCli{}

	mirror_list, err := a.Mirror_list()
	if err != nil {
		return err
	}

	ctx.Logging.Info.Println("mirror information successfully retrieved")
	fmt.Println("Mirrors:")
	utils.PrintSlice(mirror_list)

	repo_list, err := a.Repo_list()
	if err != nil {
		return err
	}

	ctx.Logging.Info.Println("repo information successfully retrieved")
	fmt.Println("Repos:")
	utils.PrintSlice(repo_list)

	return nil
}

func makeCmdDump() *commander.Command {
	cmd := &commander.Command{
		Run:       dump,
		UsageLine: "dump",
		Short:     "Dump current aptly configuration",
	}
	return cmd
}
