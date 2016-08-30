package cmd

import (
	"github.com/gonuts/commander"
	ctx "github.com/queeno/aptlify/context"
)

func Run(cmd *commander.Command, args []string, initialise bool) (returnCode int) {
	ctx.Logging.Trace.Println("starting aptlify")

	var err error

	if initialise {
		err = initContext()
		if err != nil {
			ctx.Logging.Fatal.Fatalf(err.Error())
		}
	}
	defer shutdownContext()

	err = cmd.Dispatch(args)

	if err != nil {
		ctx.Logging.Fatal.Fatalf(err.Error())
	}

	return 0

}
