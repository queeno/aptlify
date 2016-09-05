package context

import (
	"fmt"
	"github.com/queeno/aptlify/config"
	"github.com/queeno/aptlify/utils"
	"os"
	"path/filepath"
)

type AptlifyContext struct {
	config_loaded bool
	state_loaded bool
}

var Logging = utils.NewLogging()

// Creates a new configuration
func NewContext() (*AptlifyContext, error) {

	context := &AptlifyContext{
		config_loaded: false,
		state_loaded: false,
	}

	return context, nil

}

// Shutsdown the context

func ShutdownContext() error {
	return nil
}

func (context *AptlifyContext) State() *config.ConfigStruct {

	if context.state_loaded {
		return &config.State
	}

	var err error

	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
  filePath = filepath.Join(filePath, "aptlify.state")
	if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("error loading current path: %s, %s", filePath, err))
	}

	err = config.LoadConfig(filePath, &config.State)
	if os.IsNotExist(err) {
		Logging.Warning.Println(fmt.Sprintf("state file does not exist %s: %s", filePath, err))
		config.WriteConfig(filePath, config.State)
	} else if err != nil {
		Logging.Fatal.Fatalf(fmt.Sprintf("error loading state file %s, %s", filePath, err))
	}

	context.state_loaded = true

	return &config.State

}

// Load configuration and inject it into context
func (context *AptlifyContext) Config() *config.ConfigStruct {

	if context.config_loaded {
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

	context.config_loaded = true

	return &config.Config

}
