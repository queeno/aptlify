package context

import (
	"fmt"
	"github.com/queeno/aptlify/config"
	"github.com/queeno/aptlify/utils"
	"os"
	"path/filepath"
)

type AptlifyContext struct {
	configLoaded bool
	stateLoaded  bool
}

var Logging = utils.NewLogging()

// Creates a new configuration
func NewContext() (*AptlifyContext, error) {

	context := &AptlifyContext{
		configLoaded: false,
		stateLoaded:  false,
	}

	return context, nil

}

// Shuts down the context

func ShutdownContext() error {
	return nil
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

	if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("File not found!"))
	}

	context.configLoaded = true

	return &config.Config

}
