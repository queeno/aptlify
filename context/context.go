package context

import (
	"fmt"
	"github.com/queeno/aptlify/config"
	"os"
	"path/filepath"
)

type AptlifyContext struct {
	config_loaded bool
}

var Logging = utils.NewLogging()

// Creates a new configuration
func NewContext() (*AptlifyContext, error) {

	context := &AptlifyContext{
		config_loaded: false,
	}

	return context, nil

}

// Shutsdown the context

func ShutdownContext() error {
	return nil
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
