package custom_logger

import (
	"fmt"
	"os"

	"github.com/apsdehal/go-logger"
)

type CustomLogger interface {
	Infof(format string, a ...interface{})
	Debugf(format string, a ...interface{})
	Warningf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
}

func InitLogger(moduleName string) CustomLogger {
	log, err := logger.New(moduleName, 1, os.Stdout)

	if err != nil {
		fmt.Printf("error while initializing logger: %v", err)
		os.Exit(1)
	}

	return log
}
