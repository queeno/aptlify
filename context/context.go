package context

import (
	"fmt"
	"github.com/queeno/aptlify/config"
	"github.com/queeno/aptlify/utils"
	"github.com/smira/flag"
	"os"
	"path/filepath"
)

type AptlifyContext struct {
	flags, globalFlags *flag.FlagSet
	configLoaded       bool
	stateLoaded        bool
}

var Logging = utils.NewLogging()

// Creates a new configuration
func NewContext(flags *flag.FlagSet) (*AptlifyContext, error) {

	context := &AptlifyContext{
		flags:        flags,
		globalFlags:  flags,
		configLoaded: false,
		stateLoaded:  false,
	}

	return context, nil

}

// Shuts down the context

func ShutdownContext() error {
	return nil
}

// LookupOption (public) checks boolean flag
func (context *AptlifyContext) LookupOption(defaultValue bool, name string) (result bool) {
	return context.lookupOption(defaultValue, name)
}

// lookupOption (private)
func (context *AptlifyContext) lookupOption(defaultValue bool, name string) (result bool) {
	result = defaultValue
	fmt.Println(context.globalFlags)
	if context.globalFlags.IsSet(name) {
		result = context.globalFlags.Lookup(name).Value.Get().(bool)
	}
	return
}

// set context copy of flags
func (context *AptlifyContext) UpdateFlags(flags *flag.FlagSet) {
	context.flags = flags
}

// get Flags
func (context *AptlifyContext) Flags() *flag.FlagSet {
	return context.flags
}

// GlobalFlags returns flags passed to all commands
func (context *AptlifyContext) GlobalFlags() *flag.FlagSet {
	return context.globalFlags
}

func (context *AptlifyContext) WriteState(state config.ConfigStruct) {

	var err error

	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath = filepath.Join(filePath, "aptlify.state")
	if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("error loading current path: %s, %s", filePath, err))
	}

	err = config.WriteConfig(filePath, state)

	if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("error writing state file %s, %s", filePath, err))
	}

}

func (context *AptlifyContext) State() *config.ConfigStruct {

	if context.stateLoaded {
		return &config.State
	}

	var err error

	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath = filepath.Join(filePath, "aptlify.state")
	if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("error loading current path: %s, %s", filePath, err))
	}

	err = config.LoadConfig(filePath, &config.State)
	if err != nil && !os.IsNotExist(err) {
		Logging.Fatal.Fatalf(fmt.Sprintf("error loading state file %s, %s", filePath, err))
	}

	context.stateLoaded = true

	return &config.State

}

// Load configuration and inject it into context
func (context *AptlifyContext) Config() *config.ConfigStruct {

	if context.configLoaded {
		return &config.Config
	}

	var err error

	configLocation := context.globalFlags.Lookup("config").Value.String()
	if configLocation != "" {
		err = config.LoadConfig(configLocation, &config.Config)
		if err != nil {
			Logging.Fatal.Fatalf(fmt.Sprintf("error loading config file %s, %s", configLocation, err))
		}
	} else {
		filePaths := []string{
			filepath.Join(os.Getenv("HOME"), ".aptlify.conf"),
			"/etc/aptlify.conf",
		}
		for _, filePath := range filePaths {
			err = config.LoadConfig(filePath, &config.Config)
			if err == nil {
				break
			}

			if !os.IsNotExist(err) {
				Logging.Fatal.Fatalf(fmt.Sprintf("error loading config file %s, %s", filePath, err))
			}
		}
	}

	if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("File not found!"))
	}

	context.configLoaded = true

	return &config.Config

}
