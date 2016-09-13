package cmd

import (
	ctx "github.com/queeno/aptlify/context"
	"github.com/smira/commander"
)

func Run(cmd *commander.Command, args []string, initialise bool) (returnCode int) {
	ctx.Logging.Trace.Println("starting aptlify")

	var err error

	flags, args, err := cmd.ParseFlags(args)
	if err != nil {
		ctx.Logging.Fatal.Fatalf(err.Error())
	}

	if initialise {
		err = initContext(flags)
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
