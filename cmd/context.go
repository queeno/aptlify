package cmd

import (
	ctx "github.com/queeno/aptlify/context"
	"github.com/smira/flag"
)

var context *ctx.AptlifyContext

// Initialise context
func initContext(flags *flag.FlagSet) error {

	var err error

	if context != nil {
		ctx.Logging.Fatal.Fatalf("Context already initialised")
	}

	context, err = ctx.NewContext(flags)

	return err
}

func shutdownContext() {

	var err error

	if context == nil {
		ctx.Logging.Fatal.Fatalf("Shutdown context when not initialised")
	}

	err = ctx.ShutdownContext()

	if err != nil {
		ctx.Logging.Fatal.Fatalf(err.Error())
	}

}
